package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IncidentMessageV2 struct {
	// 类型，0客户留言 1华为工程师留言

	Type *int32 `json:"type,omitempty"`
	// 回复人ID

	Replier *string `json:"replier,omitempty"`
	// 留言内容

	Content *string `json:"content,omitempty"`
	// 留言id

	MessageId *string `json:"message_id,omitempty"`
	// 回复人名称

	ReplierName *string `json:"replier_name,omitempty"`
	// 创建时间

	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`
	// 是否是第一条留言

	IsFirstMessage *int32 `json:"is_first_message,omitempty"`
	// 附件列表

	AccessoryList *[]SimpleAccessoryV2 `json:"accessory_list,omitempty"`
}

func (o IncidentMessageV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IncidentMessageV2 struct{}"
	}

	return strings.Join([]string{"IncidentMessageV2", string(data)}, " ")
}
