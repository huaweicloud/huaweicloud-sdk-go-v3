package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOrUpdateRemediationConfigurationResponse Response Object
type CreateOrUpdateRemediationConfigurationResponse struct {

	// 是否为自动修正。
	Automatic *bool `json:"automatic,omitempty"`

	// 合规规则修正执行的方式。
	TargetType *string `json:"target_type,omitempty"`

	// 修正执行的目标ID。如果修正方式为fgs，则该值为函数工作流的函数urn；如果修正方式为rfs，则该值为资源编排服务的模板name与版本号，两者以/分割，如果没有指定默认V1。
	TargetId *string `json:"target_id,omitempty"`

	// 修正执行的参数。
	StaticParameter *[]RemediationStaticParameter `json:"static_parameter,omitempty"`

	ResourceParameter *RemediationResourceParameter `json:"resource_parameter,omitempty"`

	// 指定时间内修正的最大尝试次数。
	MaximumAttempts *int32 `json:"maximum_attempts,omitempty"`

	// 用于防止循环修正的时间窗口，如果在指定时间内进行了自动修正的最大尝试次数，则将资源添加至修正例外。
	RetryAttemptSeconds *int32 `json:"retry_attempt_seconds,omitempty"`

	// 合规规则修正配置的权限方式。
	AuthType *string `json:"auth_type,omitempty"`

	// 合规规则修正配置的权限信息。
	AuthValue *string `json:"auth_value,omitempty"`

	// 修正配置的创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 修正配置的更新时间。
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 创建者。
	CreatedBy      *string `json:"created_by,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateOrUpdateRemediationConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrUpdateRemediationConfigurationResponse struct{}"
	}

	return strings.Join([]string{"CreateOrUpdateRemediationConfigurationResponse", string(data)}, " ")
}
