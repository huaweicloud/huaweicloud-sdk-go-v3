package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Policy 身份策略。
type Policy struct {
	PolicyType *PolicyType `json:"policy_type"`

	// 身份策略名称，长度为1到128个字符，只包含字母、数字、\"_\"、\"+\"、\"=\"、\".\"、\"@\"和\"-\"的字符串。
	PolicyName string `json:"policy_name"`

	// 身份策略ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	PolicyId string `json:"policy_id"`

	// 统一资源名称。
	Urn string `json:"urn"`

	// 资源路径，默认为空串。由若干段字符串拼接而成，每段先包含一个或多个字母、数字、\".\"、\",\"、\"+\"、\"@\"、\"=\"、\"_\"或\"-\"，并以\"/\"结尾，例如\"foo/bar/\"。
	Path string `json:"path"`

	// 默认版本号。
	DefaultVersionId string `json:"default_version_id"`

	// 附加了本身份策略的实体数量。
	AttachmentCount int32 `json:"attachment_count"`

	// 身份策略描述。
	Description *string `json:"description,omitempty"`

	// 身份策略创建时间。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 身份策略默认版本最近一次的更新时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`
}

func (o Policy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Policy struct{}"
	}

	return strings.Join([]string{"Policy", string(data)}, " ")
}
