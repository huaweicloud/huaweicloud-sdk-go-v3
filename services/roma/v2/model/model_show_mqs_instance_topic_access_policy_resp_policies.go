package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowMqsInstanceTopicAccessPolicyRespPolicies struct {

	// 是否为创建topic时所选择的应用。
	Owner *bool `json:"owner,omitempty" xml:"owner"`

	// 应用ID。
	UserName *string `json:"user_name,omitempty" xml:"user_name"`

	// 权限类型。   - all：发布+订阅   - pub：发布   - sub：订阅
	AccessPolicy *string `json:"access_policy,omitempty" xml:"access_policy"`

	// 应用名称。
	AppName *string `json:"app_name,omitempty" xml:"app_name"`

	// 权限类型对应的标签。
	Tag *string `json:"tag,omitempty" xml:"tag"`
}

func (o ShowMqsInstanceTopicAccessPolicyRespPolicies) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMqsInstanceTopicAccessPolicyRespPolicies struct{}"
	}

	return strings.Join([]string{"ShowMqsInstanceTopicAccessPolicyRespPolicies", string(data)}, " ")
}
