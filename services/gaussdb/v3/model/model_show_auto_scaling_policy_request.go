package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutoScalingPolicyRequest Request Object
type ShowAutoScalingPolicyRequest struct {

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`

	// 内容类型。  取值：application/json。
	ContentType string `json:"Content-Type"`

	// 实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`
}

func (o ShowAutoScalingPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoScalingPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowAutoScalingPolicyRequest", string(data)}, " ")
}
