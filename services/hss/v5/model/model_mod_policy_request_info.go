package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModPolicyRequestInfo 修改策略请求
type ModPolicyRequestInfo struct {

	// 策略名称
	PolicyName string `json:"policy_name"`

	// 策略ID
	PolicyId string `json:"policy_id"`

	// 模板ID
	TemplateId string `json:"template_id"`

	// 策略描述
	PolicyDescription string `json:"policy_description"`

	Content *PolicyContent `json:"content"`

	// 白名单镜像
	WhiteImageList []WhiteImageInfo `json:"white_image_list"`

	// 防护的资源信息
	Resources []PolicyProtectScope `json:"resources"`

	// 策略参数值
	Parameters string `json:"parameters"`
}

func (o ModPolicyRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModPolicyRequestInfo struct{}"
	}

	return strings.Join([]string{"ModPolicyRequestInfo", string(data)}, " ")
}
