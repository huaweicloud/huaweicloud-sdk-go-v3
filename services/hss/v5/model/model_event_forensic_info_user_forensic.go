package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventForensicInfoUserForensic **参数解释**： 用户调查取证信息 **取值范围**： 不涉及
type EventForensicInfoUserForensic struct {

	// **参数解释**： 用户uid **取值范围**： 不涉及
	UserId *int32 `json:"user_id,omitempty"`

	// **参数解释**： 用户gid **取值范围**： 不涉及
	UserGid *int32 `json:"user_gid,omitempty"`

	// **参数解释**： 用户名称 **取值范围**： 不涉及
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**： 用户组名称 **取值范围**： 不涉及
	UserGroupName *string `json:"user_group_name,omitempty"`

	// **参数解释**： 用户home目录 **取值范围**： 不涉及
	UserHomeDir *string `json:"user_home_dir,omitempty"`

	// **参数解释**： 用户登录ip **取值范围**： 不涉及
	LoginIp *string `json:"login_ip,omitempty"`

	// **参数解释**： 登录的服务类型 **取值范围**： 不涉及
	ServiceType *string `json:"service_type,omitempty"`

	// **参数解释**： 登录服务端口 **取值范围**： 不涉及
	ServicePort *int32 `json:"service_port,omitempty"`

	// **参数解释**： 登录方式 **取值范围**： 不涉及
	LoginMode *int32 `json:"login_mode,omitempty"`

	// **参数解释**： 用户最后一次登录时间 **取值范围**： 不涉及
	LoginLastTime *int64 `json:"login_last_time,omitempty"`

	// **参数解释**： 用户登录失败次数 **取值范围**： 不涉及
	LoginFailCount *int32 `json:"login_fail_count,omitempty"`
}

func (o EventForensicInfoUserForensic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventForensicInfoUserForensic struct{}"
	}

	return strings.Join([]string{"EventForensicInfoUserForensic", string(data)}, " ")
}
