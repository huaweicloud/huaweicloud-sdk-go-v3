package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClusterResponse Response Object
type CreateClusterResponse struct {

	// 集群名称
	Name *string `json:"name,omitempty"`

	// 集群ID
	Id *string `json:"id,omitempty"`

	Task *Task `json:"task,omitempty"`

	Datastore *Datastore `json:"datastore,omitempty"`

	// 集群的节点信息
	Instances      *[]ClusterInstance `json:"instances,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o CreateClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterResponse struct{}"
	}

	return strings.Join([]string{"CreateClusterResponse", string(data)}, " ")
}
