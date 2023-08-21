package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplyWafPolicyResponse Response Object
type ApplyWafPolicyResponse struct {

	// 防护策略id
	Id *string `json:"id,omitempty"`

	// 防护策略名
	Name *string `json:"name,omitempty"`

	Action *WafPolicyAction `json:"action,omitempty"`

	Options *WafPolicyOption `json:"options,omitempty"`

	// 防护等级
	Level *int32 `json:"level,omitempty"`

	// 精准防护中的检测模式
	FullDetection *bool `json:"full_detection,omitempty"`

	RobotAction *WafPolicyAction `json:"robot_action,omitempty"`

	// 防护域名的信息
	BindHost *[]WafPolicyBindHost `json:"bind_host,omitempty"`

	// 创建防护策略的时间
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 扩展字段
	Extend         map[string]string `json:"extend,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ApplyWafPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyWafPolicyResponse struct{}"
	}

	return strings.Join([]string{"ApplyWafPolicyResponse", string(data)}, " ")
}
