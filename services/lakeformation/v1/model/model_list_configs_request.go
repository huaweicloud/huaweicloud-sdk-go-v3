package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConfigsRequest Request Object
type ListConfigsRequest struct {

	// LakeFormation实例ID。创建实例时自动生成。例如：2180518f-42b8-4947-b20b-adfc53981a25。
	InstanceId string `json:"instance_id"`

	// 查询返回条数。默认值为1000。最小值为0，最大值为1000。
	Limit *int32 `json:"limit,omitempty"`

	// 查询的起始记录ID。最小长度为0，最大长度为256。
	Marker *string `json:"marker,omitempty"`
}

func (o ListConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigsRequest struct{}"
	}

	return strings.Join([]string{"ListConfigsRequest", string(data)}, " ")
}
