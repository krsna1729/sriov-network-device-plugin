package network

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NetAnnonKey is the key to the annotation map
const NetAnnonKey = "kubernetes.v1.cni.cncf.io/networks"

// Network represents the CRD proposed by NW Plumbing WG
// copied from github.com/intel/multus-cni/blob/v2.0/types/types.go
type Network struct {
	metav1.TypeMeta `json:",inline"`
	// Note that ObjectMeta is mandatory, as an object
	// name is required
	Metadata metav1.ObjectMeta `json:"metadata,omitempty" description:"standard object metadata"`

	// Specification describing how to invoke a CNI plugin to
	// add or remove network attachments for a Pod.
	// In the absence of valid keys in a Spec, the runtime (or
	// meta-plugin) should load and execute a CNI .configlist
	// or .config (in that order) file on-disk whose JSON
	// “name” key matches this Network object’s name.
	// +optional
	Spec Spec `json:"spec"`
}

// Spec represents the CNI information
type Spec struct {
	// Config contains a standard JSON-encoded CNI configuration
	// or configuration list which defines the plugin chain to
	// execute.  If present, this key takes precedence over
	// ‘Plugin’.
	// +optional
	Config string `json:"config"`

	// Plugin contains the name of a CNI plugin on-disk in a
	// runtime-defined path (eg /opt/cni/bin and/or other paths.
	// This plugin should be executed with a basic CNI JSON
	// configuration on stdin containing the Network object
	// name and the plugin:
	//   { “cniVersion”: “0.3.1”, “type”: <Plugin>, “name”: <Network.Name> }
	// and any additional “runtimeConfig” field per the
	// CNI specification and conventions.
	// +optional
	Plugin string `json:"plugin"`
}
