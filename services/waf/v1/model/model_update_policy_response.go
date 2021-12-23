package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdatePolicyResponse struct {
	// 防护策略id

	Id *string `json:"id,omitempty"`
	// 防护策略名

	Name *string `json:"name,omitempty"`
	// 防护等级

	Level *int32 `json:"level,omitempty"`

	Action *PolicyAction `json:"action,omitempty"`

	Options *PolicyOption `json:"options,omitempty"`
	// 防护域名的信息

	Hosts *[]string `json:"hosts,omitempty"`
	// 扩展字段

	Extend map[string]string `json:"extend,omitempty"`
	// 创建防护策略的时间

	Timestamp *int64 `json:"timestamp,omitempty"`
	// 精准防护中的检测模式

	FullDetection *bool `json:"full_detection,omitempty"`
	// 防护域名的信息

	BindHost       *[]BindHost `json:"bind_host,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o UpdatePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdatePolicyResponse", string(data)}, " ")
}
