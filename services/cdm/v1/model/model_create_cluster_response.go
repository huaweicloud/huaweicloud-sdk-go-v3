package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateClusterResponse struct {

	// 集群名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 集群ID
	Id *string `json:"id,omitempty" xml:"id"`

	Task *Task `json:"task,omitempty" xml:"task"`

	Datastore *Datastore `json:"datastore,omitempty" xml:"datastore"`

	// 集群的节点信息
	Instances      *[]ClusterInstance `json:"instances,omitempty" xml:"instances"`
	HttpStatusCode int                `json:"-"`
}

func (o CreateClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterResponse struct{}"
	}

	return strings.Join([]string{"CreateClusterResponse", string(data)}, " ")
}
