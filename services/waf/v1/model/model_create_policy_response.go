package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreatePolicyResponse struct {

	// 防护策略id
	Id *string `json:"id,omitempty"`

	// 防护策略名
	Name *string `json:"name,omitempty"`

	// 防护等级
	Level *int32 `json:"level,omitempty"`

	Action *PolicyAction `json:"action,omitempty"`

	Options *PolicyOption `json:"options,omitempty"`

	// 精准防护中的检测模式
	FullDetection *bool `json:"full_detection,omitempty"`

	// 防护的网站id
	Hosts *[]string `json:"hosts,omitempty"`

	// 防护的网站信息
	BindHost *[]BindHost `json:"bind_host,omitempty"`

	// 创建防护策略的时间
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 扩展字段
	Extend         *interface{} `json:"extend,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o CreatePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePolicyResponse struct{}"
	}

	return strings.Join([]string{"CreatePolicyResponse", string(data)}, " ")
}
