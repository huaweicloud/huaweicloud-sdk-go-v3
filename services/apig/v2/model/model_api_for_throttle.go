package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApiForThrottle struct {

	// API的认证方式
	AuthType *string `json:"auth_type,omitempty" xml:"auth_type"`

	// API所属分组的名称
	GroupName *string `json:"group_name,omitempty" xml:"group_name"`

	// API的发布记录编号
	PublishId *string `json:"publish_id,omitempty" xml:"publish_id"`

	// 与流控策略的绑定关系编号
	ThrottleApplyId *string `json:"throttle_apply_id,omitempty" xml:"throttle_apply_id"`

	// 已绑定的流控策略的绑定时间
	ApplyTime *sdktime.SdkTime `json:"apply_time,omitempty" xml:"apply_time"`

	// API描述
	Remark *string `json:"remark,omitempty" xml:"remark"`

	// 发布的环境id
	RunEnvId *string `json:"run_env_id,omitempty" xml:"run_env_id"`

	// API类型
	Type *int32 `json:"type,omitempty" xml:"type"`

	// 绑定的流控策略名称
	ThrottleName *string `json:"throttle_name,omitempty" xml:"throttle_name"`

	// API的访问地址
	ReqUri *string `json:"req_uri,omitempty" xml:"req_uri"`

	// 发布的环境名
	RunEnvName *string `json:"run_env_name,omitempty" xml:"run_env_name"`

	// API所属分组的编号
	GroupId *string `json:"group_id,omitempty" xml:"group_id"`

	// API名称
	Name *string `json:"name,omitempty" xml:"name"`

	// API编号
	Id *string `json:"id,omitempty" xml:"id"`
}

func (o ApiForThrottle) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiForThrottle struct{}"
	}

	return strings.Join([]string{"ApiForThrottle", string(data)}, " ")
}
