package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// ZookeeperSpec defines the desired state of Zookeeper
type ZookeeperSpec struct {
	LanIP       string `json:"lanIP"`
	ClusterID   string `json:"clusterID"`
	ClusterName string `json:"clusterName"`
	// +optional
	Import bool `json:"import"`
	// +optional
	SDU string `json:"sdu"`
	// +optional
	NodeID string `json:"nodeID"`
	// +optional
	AZ string `json:"az"`
	// +optional
	Env string `json:"env"`
	// +optional
	IDC string `json:"idc"`
	// +optional
	CID string `json:"cid"`
	// +optional
	TOCCluster string `json:"tocCluster"`
	// +optional
	Kernel string `json:"kernel"`
	// +optional
	ServerID uint64 `json:"serverID"`
	// +optional
	ProductID uint64 `json:"productID"`
	// +optional
	ServiceTag string `json:"serviceTag"`
	// +optional
	ClusterInfo string `json:"clusterInfo"`
	// +optional
	ZookeeperVersion string `json:"zookeeperVersion"`
	// +optional
	ZookeeperTag string `json:"zookeeperTag"`
	// +optional
	DockerVersion string `json:"dockerVersion"`
}

// ZookeeperStatus defines the observed state of Zookeeper.
// It should always be reconstructable from the state of the cluster and/or outside world.
type ZookeeperStatus struct {
	// OperationTimestamp is the server time when this PhysicalMachine was started
	// +optional
	OperationTimestamp metav1.Time `json:"operationTimestamp,omitempty"`
	State              string      `json:"state,omitempty"`
	Reason             string      `json:"reason,omitempty"`
	// LastTransitionTime is the timestamp of the last PhysicalMachine transition
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Zookeeper is the Schema for the zookeepers API
// +k8s:openapi-gen=true
type Zookeeper struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ZookeeperSpec   `json:"spec,omitempty"`
	Status ZookeeperStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ZookeeperList contains a list of Zookeeper
type ZookeeperList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Zookeeper `json:"items"`
}
