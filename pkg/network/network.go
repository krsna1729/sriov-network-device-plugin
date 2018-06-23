package network

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"k8s.io/client-go/kubernetes"
)

func parsePodNetworkObjectName(podnetwork string) (string, string, string, error) {
	var netNsName string
	var netIfName string
	var networkName string

	slashItems := strings.Split(podnetwork, "/")
	if len(slashItems) == 2 {
		netNsName = strings.TrimSpace(slashItems[0])
		networkName = slashItems[1]
	} else if len(slashItems) == 1 {
		networkName = slashItems[0]
	} else {
		return "", "", "", fmt.Errorf("Invalid network object (failed at '/')")
	}

	atItems := strings.Split(networkName, "@")
	networkName = strings.TrimSpace(atItems[0])
	if len(atItems) == 2 {
		netIfName = strings.TrimSpace(atItems[1])
	} else if len(atItems) != 1 {
		return "", "", "", fmt.Errorf("Invalid network object (failed at '@')")
	}

	// Check and see if each item matches the specification for valid attachment name.
	// "Valid attachment names must be comprised of units of the DNS-1123 label format"
	// [a-z0-9]([-a-z0-9]*[a-z0-9])?
	// And we allow at (@), and forward slash (/) (units separated by commas)
	// It must start and end alphanumerically.
	allItems := []string{netNsName, networkName, netIfName}
	for i := range allItems {
		matched, _ := regexp.MatchString("^[a-z0-9]([-a-z0-9]*[a-z0-9])?$", allItems[i])
		if !matched && len([]rune(allItems[i])) > 0 {
			return "", "", "", fmt.Errorf(fmt.Sprintf("Failed to parse: one or more items did not match comma-delimited format (must consist of lower case alphanumeric characters). Must start and end with an alphanumeric character), mismatch @ '%v'", allItems[i]))
		}
	}

	return netNsName, networkName, netIfName, nil
}

func parsePodNetworkObject(podnetwork string) ([]map[string]interface{}, error) {
	var podNet []map[string]interface{}

	if podnetwork == "" {
		return nil, fmt.Errorf("parsePodNetworkObject: pod annotation not having \"network\" as key, refer Multus README.md for the usage guide")
	}

	// Parse the podnetwork string, and assume it is JSON.
	if err := json.Unmarshal([]byte(podnetwork), &podNet); err != nil {

		// If JSON doesn't parse, assume comma-delimited.
		commaItems := strings.Split(podnetwork, ",")

		// Build a map from the comma delimited items.
		for i := range commaItems {

			// Remove leading and trailing whitespace.
			commaItems[i] = strings.TrimSpace(commaItems[i])

			// Parse network name (i.e. <namespace>/<network name>@<ifname>)
			netNsName, networkName, netIfName, err := parsePodNetworkObjectName(commaItems[i])
			if err != nil {
				return nil, fmt.Errorf("parsePodNetworkObject: %v", err)
			}
			m := make(map[string]interface{})
			m["name"] = networkName
			if netNsName != "" {
				m["namespace"] = netNsName
			}
			if netIfName != "" {
				m["interfaceRequest"] = netIfName
			}

			podNet = append(podNet, m)
		}
	}

	return podNet, nil
}

func getnetplugin(client *kubernetes.Clientset, networkinfo map[string]interface{}) (*Network, error) {

	networkname := networkinfo["name"].(string)
	if networkname == "" {
		return nil, fmt.Errorf("getnetplugin: network name can't be empty")
	}

	netNsName := "default"
	if networkinfo["namespace"] != nil {
		netNsName = networkinfo["namespace"].(string)
	}

	tprclient := fmt.Sprintf("/apis/kubernetes.cni.cncf.io/v1/namespaces/%s/networks/%s", netNsName, networkname)

	netobjdata, err := client.ExtensionsV1beta1().RESTClient().Get().AbsPath(tprclient).DoRaw()
	if err != nil {
		return nil, fmt.Errorf("getnetplugin: failed to get CRD (result: %s), refer Multus README.md for the usage guide: %v", netobjdata, err)
	}
	netobj := Network{}
	if err := json.Unmarshal(netobjdata, &netobj); err != nil {
		return nil, fmt.Errorf("getnetplugin: failed to get the netplugin data: %v", err)
	}

	return &netobj, nil
}

func getPodNetworkObj(client *kubernetes.Clientset, netObjs []map[string]interface{}) ([]*Network, error) {
	var nets []*Network
	for _, no := range netObjs {
		net, err := getnetplugin(client, no)
		if err != nil {
			return nil, err
		}
		nets = append(nets, net)

	}
	return nets, nil
}

// Getk8sNets returns configs of all networks requested in the annotations
func Getk8sNets(k8sclient *kubernetes.Clientset, netAnnot string) ([]*Network, error) {

	if len(netAnnot) == 0 {
		return nil, fmt.Errorf("No networks found")
	}

	netObjs, err := parsePodNetworkObject(netAnnot)
	if err != nil {
		return nil, err
	}

	multusDelegates, err := getPodNetworkObj(k8sclient, netObjs)
	if err != nil {
		return nil, err
	}
	return multusDelegates, nil
}
