package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSmartChatJobRequest Request Object
type ListSmartChatJobRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 第三方用户ID。不允许输入中文。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	// 智能交互对话ID，获取方法请参考[创建智能交互对话直播间](CreateSmartChatRoom.xml)。
	RoomId string `json:"room_id"`

	// 任务状态，默认所有状态。  可多个状态查询，使用英文逗号分隔。  如state=CREATING,PUBLISHED
	State *string `json:"state,omitempty"`

	// 偏移量，表示从此偏移量开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListSmartChatJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSmartChatJobRequest struct{}"
	}

	return strings.Join([]string{"ListSmartChatJobRequest", string(data)}, " ")
}
