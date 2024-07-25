package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TmsResourceResp struct {

	// 实例编号
	ResourceId *string `json:"resource_id,omitempty"`

	// 实例详细描述。暂不支持
	ResouceDetail *string `json:"resouce_detail,omitempty"`

	// 实例名称
	ResourceName *string `json:"resource_name,omitempty"`

	// 实例绑定的标签列表
	Tags *[]TmsKeyValue `json:"tags,omitempty"`
}

func (o TmsResourceResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TmsResourceResp struct{}"
	}

	return strings.Join([]string{"TmsResourceResp", string(data)}, " ")
}
