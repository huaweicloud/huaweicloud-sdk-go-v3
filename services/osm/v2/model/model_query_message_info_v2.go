package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryMessageInfoV2 struct {

	// 类型，0客户留言 1华为工程师留言
	Type *int32 `json:"type,omitempty" xml:"type"`

	// 回复人类型，0客户留言 1华为工程师留言 2第三方留言
	ReplierType *int32 `json:"replier_type,omitempty" xml:"replier_type"`

	// 回复人id
	Replier *string `json:"replier,omitempty" xml:"replier"`

	// 留言内容
	Content *string `json:"content,omitempty" xml:"content"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty" xml:"create_time"`

	// 回复人名称
	ReplierName *string `json:"replier_name,omitempty" xml:"replier_name"`

	// 是否是第一条留言
	IsFirstMessage *int32 `json:"is_first_message,omitempty" xml:"is_first_message"`

	// 子用户类型
	IamUserType *int32 `json:"iam_user_type,omitempty" xml:"iam_user_type"`

	// 附件列表
	AccessoryList *[]SimpleAccessoryV2 `json:"accessory_list,omitempty" xml:"accessory_list"`
}

func (o QueryMessageInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryMessageInfoV2 struct{}"
	}

	return strings.Join([]string{"QueryMessageInfoV2", string(data)}, " ")
}
