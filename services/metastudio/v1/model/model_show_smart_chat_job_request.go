package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSmartChatJobRequest Request Object
type ShowSmartChatJobRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 第三方用户ID。不允许输入中文。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	// 智能交互对话ID，获取方法请参考[[创建智能交互对话直播间](https://support.huaweicloud.com/api-metastudio/CreateSmartChatRoom.html)](tag:hc,hk)[“创建智能交互对话直播间”](tag:cmcc)。
	RoomId string `json:"room_id"`

	// 任务ID。
	JobId string `json:"job_id"`
}

func (o ShowSmartChatJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSmartChatJobRequest struct{}"
	}

	return strings.Join([]string{"ShowSmartChatJobRequest", string(data)}, " ")
}
