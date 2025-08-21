package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVpcepConnectionRequest Request Object
type ShowVpcepConnectionRequest struct {

	// 指定待查询的集群ID。
	ClusterId string `json:"cluster_id"`

	// 指定查询起始值，默认值为1，即从第1个任务开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 指定查询个数，默认值为10，即一次查询10个任务信息。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowVpcepConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVpcepConnectionRequest struct{}"
	}

	return strings.Join([]string{"ShowVpcepConnectionRequest", string(data)}, " ")
}
