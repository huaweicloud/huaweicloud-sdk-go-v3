package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateExtendClusterResponse struct {
	// 集群ID。

	Id *string `json:"id,omitempty"`
	// 扩容实例列表。

	Instances      *[]ExtendClusterInstancesResp `json:"instances,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o UpdateExtendClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateExtendClusterResponse struct{}"
	}

	return strings.Join([]string{"UpdateExtendClusterResponse", string(data)}, " ")
}
