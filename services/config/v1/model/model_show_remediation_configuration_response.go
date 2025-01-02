package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRemediationConfigurationResponse Response Object
type ShowRemediationConfigurationResponse struct {

	// 是否为自动修正。
	Automatic *bool `json:"automatic,omitempty"`

	// 合规规则修正执行的方式。
	TargetType *string `json:"target_type,omitempty"`

	// 修正执行的目标ID。如果修正方式为fgs，则该值为函数工作流的函数urn；如果修正方式为rfs，则该值为资源编排服务的模板name与版本号，两者以/分割，如果没有指定默认V1。
	TargetId *string `json:"target_id,omitempty"`

	// 修正执行的目标的regionId。如果修正方式为RFS，该字段为空则Config服务会默认配置北京四（中国站）或香港一（国际站）的regionId；如果修正方式为FGS，该字段为空则Config服务会根据实例urn自动配置。指定target_project_id字段则该字段必选。
	TargetRegionId *string `json:"target_region_id,omitempty"`

	// 修正执行的目标的projectId。如果修正方式为RFS，该字段为空则Config服务会默认配置北京四（中国站）或香港一（国际站）的主projectId；如果修正方式为FGS，该字段为空则Config服务会根据实例urn自动配置。指定target_region_id字段则该字段必选。
	TargetProjectId *string `json:"target_project_id,omitempty"`

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

func (o ShowRemediationConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRemediationConfigurationResponse struct{}"
	}

	return strings.Join([]string{"ShowRemediationConfigurationResponse", string(data)}, " ")
}
