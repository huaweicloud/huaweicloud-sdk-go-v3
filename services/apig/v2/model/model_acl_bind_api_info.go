package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AclBindApiInfo struct {

	// API编号
	ApiId *string `json:"api_id,omitempty" xml:"api_id"`

	// API名称
	ApiName *string `json:"api_name,omitempty" xml:"api_name"`

	// API类型
	ApiType *int64 `json:"api_type,omitempty" xml:"api_type"`

	// API的描述信息
	ApiRemark *string `json:"api_remark,omitempty" xml:"api_remark"`

	// 生效的环境编号
	EnvId *string `json:"env_id,omitempty" xml:"env_id"`

	// 生效的环境名称
	EnvName *string `json:"env_name,omitempty" xml:"env_name"`

	// 绑定关系编号
	BindId *string `json:"bind_id,omitempty" xml:"bind_id"`

	// API分组名称
	GroupName *string `json:"group_name,omitempty" xml:"group_name"`

	// 绑定时间
	BindTime *sdktime.SdkTime `json:"bind_time,omitempty" xml:"bind_time"`

	// API发布记录编号
	PublishId *string `json:"publish_id,omitempty" xml:"publish_id"`
}

func (o AclBindApiInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AclBindApiInfo struct{}"
	}

	return strings.Join([]string{"AclBindApiInfo", string(data)}, " ")
}
