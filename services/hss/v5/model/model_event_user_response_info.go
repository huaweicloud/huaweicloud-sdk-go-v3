package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventUserResponseInfo 用户信息
type EventUserResponseInfo struct {

	// **参数解释**： 用户uid **取值范围**： 最小值0，最大值2147483647
	UserId *int32 `json:"user_id,omitempty"`

	// **参数解释**： 用户gid **取值范围**： 最小值0，最大值2147483647
	UserGid *int32 `json:"user_gid,omitempty"`

	// **参数解释**： 用户名称 **取值范围**： 字符长度1-256位
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**： 用户组名称 **取值范围**： 字符长度1-256位
	UserGroupName *string `json:"user_group_name,omitempty"`

	// **参数解释**： 用户home目录 **取值范围**： 字符长度1-256位
	UserHomeDir *string `json:"user_home_dir,omitempty"`

	// **参数解释**： 用户登录IP **取值范围**： 字符长度1-256位
	LoginIp *string `json:"login_ip,omitempty"`

	// **参数解释**： 服务类型 **取值范围**： - system：系统 - mysql：数据库 - redis：Redis
	ServiceType *string `json:"service_type,omitempty"`

	// **参数解释**： 登录服务端口 **取值范围**： 最小值0，最大值2147483647
	ServicePort *int32 `json:"service_port,omitempty"`

	// **参数解释**： 登录方式 **取值范围**： 最小值0，最大值2147483647
	LoginMode *int32 `json:"login_mode,omitempty"`

	// **参数解释**： 用户最后一次登录时间 **取值范围**： 最小值0，最大值9223372036854775807
	LoginLastTime *int64 `json:"login_last_time,omitempty"`

	// **参数解释**： 用户登录失败次数 **取值范围**： 最小值0，最大值2147483647
	LoginFailCount *int32 `json:"login_fail_count,omitempty"`

	// **参数解释**： 口令hash **取值范围**： 字符长度1-256位
	PwdHash *string `json:"pwd_hash,omitempty"`

	// **参数解释**： 匿名化处理后的口令 **取值范围**： 字符长度1-256位
	PwdWithFuzzing *string `json:"pwd_with_fuzzing,omitempty"`

	// **参数解释**： 密码使用的天数 **取值范围**： 最小值0，最大值2147483647
	PwdUsedDays *int32 `json:"pwd_used_days,omitempty"`

	// **参数解释**： 口令的最短有效期限 **取值范围**： 最小值0，最大值2147483647
	PwdMinDays *int32 `json:"pwd_min_days,omitempty"`

	// **参数解释**： 口令的最长有效期限 **取值范围**： 最小值0，最大值2147483647
	PwdMaxDays *int32 `json:"pwd_max_days,omitempty"`

	// **参数解释**： 口令无效时提前告警天数 **取值范围**： 最小值0，最大值2147483647
	PwdWarnLeftDays *int32 `json:"pwd_warn_left_days,omitempty"`
}

func (o EventUserResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventUserResponseInfo struct{}"
	}

	return strings.Join([]string{"EventUserResponseInfo", string(data)}, " ")
}
