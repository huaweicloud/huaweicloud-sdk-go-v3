package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 用户信息
type EventUserResponseInfo struct {

	// 用户uid
	UserId *int32 `json:"user_id,omitempty" xml:"user_id"`

	// 用户gid
	UserGid *int32 `json:"user_gid,omitempty" xml:"user_gid"`

	// 用户名称
	UserName *string `json:"user_name,omitempty" xml:"user_name"`

	// 用户组名称
	UserGroupName *string `json:"user_group_name,omitempty" xml:"user_group_name"`

	// 用户home目录
	UserHomeDir *string `json:"user_home_dir,omitempty" xml:"user_home_dir"`

	// 用户登录ip
	LoginIp *string `json:"login_ip,omitempty" xml:"login_ip"`

	// 登录的服务类型
	ServiceType *string `json:"service_type,omitempty" xml:"service_type"`

	// 登录服务端口
	ServicePort *int32 `json:"service_port,omitempty" xml:"service_port"`

	// 登录方式
	LoginMode *int32 `json:"login_mode,omitempty" xml:"login_mode"`

	// 用户最后一次登录时间
	LoginLastTime *int64 `json:"login_last_time,omitempty" xml:"login_last_time"`

	// 用户登录失败次数
	LoginFailCount *int32 `json:"login_fail_count,omitempty" xml:"login_fail_count"`

	// 口令hash
	PwdHash *string `json:"pwd_hash,omitempty" xml:"pwd_hash"`

	// 匿名化处理后的口令
	PwdWithFuzzing *string `json:"pwd_with_fuzzing,omitempty" xml:"pwd_with_fuzzing"`

	// 密码使用的天数
	PwdUsedDays *int32 `json:"pwd_used_days,omitempty" xml:"pwd_used_days"`

	// 口令的最短有效期限
	PwdMinDays *int32 `json:"pwd_min_days,omitempty" xml:"pwd_min_days"`

	// 口令的最长有效期限
	PwdMaxDays *int32 `json:"pwd_max_days,omitempty" xml:"pwd_max_days"`

	// 口令无效时提前告警天数
	PwdWarnLeftDays *int32 `json:"pwd_warn_left_days,omitempty" xml:"pwd_warn_left_days"`
}

func (o EventUserResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventUserResponseInfo struct{}"
	}

	return strings.Join([]string{"EventUserResponseInfo", string(data)}, " ")
}
