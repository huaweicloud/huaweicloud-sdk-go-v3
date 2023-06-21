package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 新建智能信息基础版模板的请求消息体。
type CreateVmsTemplateRequestBody struct {

	// 智能信息基础版模板有效期。单位：天，必须取整，最长9999天。
	ExpirationTime string `json:"expiration_time"`

	// 智能信息基础版模板名称。模板的别名，用来帮助记忆。最大不超过100个字，若使用中文需经过UTF-8编码。
	TplName string `json:"tpl_name"`

	// 智能信息基础版模板主题，最大不超过20个字，若使用中文需经过UTF-8 编码，主题不能包含“【】”，否则审核会不通过。
	Title string `json:"title"`

	//  模板资源列表，由按顺序排列的资源组成，资源类型支持文本、图片、音频、视频。  > 资源在JSON数组中的顺序将决定其在手机上的显示顺序，数组大小不能超过10。
	Reslist []ResourceInfo `json:"reslist"`

	// 智能信息基础版模板备注信息，用于填写对模板审核的期望或要求，最大不超过200个字。例如：希望这个模板绑定的通道类型是三网合一通道，默认优先绑定三网合一通道。
	Remarks *string `json:"remarks,omitempty"`

	// 用于发送提醒短信，将智能信息基础版模板审批的结果和意见及时通知到指定对象，目前仅支持11位纯数字国内手机号码，多个号码用英文逗号分隔，最多支持5个手机号。
	Reminders *string `json:"reminders,omitempty"`

	// 客户系统回调URL，可用于通知对端模板审核状态信息。  > 接口规格需参照定义智能信息基础版模板状态回执完成实现。
	Callbackurl *string `json:"callbackurl,omitempty"`
}

func (o CreateVmsTemplateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVmsTemplateRequestBody struct{}"
	}

	return strings.Join([]string{"CreateVmsTemplateRequestBody", string(data)}, " ")
}
