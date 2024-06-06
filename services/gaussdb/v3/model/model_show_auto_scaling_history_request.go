package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutoScalingHistoryRequest Request Object
type ShowAutoScalingHistoryRequest struct {

	// 内容类型。  取值：application/json。
	ContentType string `json:"Content-Type"`

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 查询记录数。默认为20。
	Limit *string `json:"limit,omitempty"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *string `json:"offset,omitempty"`
}

func (o ShowAutoScalingHistoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoScalingHistoryRequest struct{}"
	}

	return strings.Join([]string{"ShowAutoScalingHistoryRequest", string(data)}, " ")
}
