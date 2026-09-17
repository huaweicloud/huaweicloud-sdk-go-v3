package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSqlLimitingInfoResponse Response Object
type ShowSqlLimitingInfoResponse struct {

	// 实例能否使用SQL限流功能
	CanUse *bool `json:"can_use,omitempty"`

	// 是否大小写敏感
	CaseSensitive *bool `json:"case_sensitive,omitempty"`

	// 是否支持展示过期
	Expire *bool `json:"expire,omitempty"`

	// 当canUse为False时展示错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 实例类型
	InstanceType *string `json:"instance_type,omitempty"`

	// 实例详细版本号
	InstanceDetailVersion *string `json:"instance_detail_version,omitempty"`

	// 只读实例是否可以添加、删除限流规则
	CanReadonlySetRule *bool `json:"can_readonly_set_rule,omitempty"`

	// 当canReadonlySetRule为False时展示的提示
	ReadonlySetRuleMsg *string `json:"readonly_set_rule_msg,omitempty"`

	// 最大可用SQL限流规则数
	MaxRuleLimit *int32 `json:"max_rule_limit,omitempty"`

	// 是否支持添加insert类型sql
	CanAddInsertType *bool `json:"can_add_insert_type,omitempty"`

	// 实例能否使用关键字自治限流功能
	SupportKeyStr  *bool `json:"support_key_str,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowSqlLimitingInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlLimitingInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowSqlLimitingInfoResponse", string(data)}, " ")
}
