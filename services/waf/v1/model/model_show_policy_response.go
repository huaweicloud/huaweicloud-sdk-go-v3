package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowPolicyResponse struct {

	// 防护策略id
	Id *string `json:"id,omitempty"`

	// 防护策略名
	Name *string `json:"name,omitempty"`

	Action *PolicyAction `json:"action,omitempty"`

	Options *PolicyOption `json:"options,omitempty"`

	// 防护等级
	Level *int32 `json:"level,omitempty"`

	// 精准防护中的检测模式
	FullDetection *bool `json:"full_detection,omitempty"`

	// 防护域名的信息
	BindHost *[]BindHost `json:"bind_host,omitempty"`

	// 创建防护策略的时间
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 扩展字段
	Extend         map[string]string `json:"extend,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowPolicyResponse", string(data)}, " ")
}
