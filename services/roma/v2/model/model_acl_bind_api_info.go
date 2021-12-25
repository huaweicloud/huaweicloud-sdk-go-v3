package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AclBindApiInfo struct {
	// API编号

	ApiId *string `json:"api_id,omitempty"`
	// API名称

	ApiName *string `json:"api_name,omitempty"`
	// API分组名称

	GroupName *string `json:"group_name,omitempty"`
	// API类型

	ApiType *int32 `json:"api_type,omitempty"`
	// API的描述信息

	ApiRemark *string `json:"api_remark,omitempty"`
	// 生效的环境编号

	EnvId *string `json:"env_id,omitempty"`
	// 生效的环境名称

	EnvName *string `json:"env_name,omitempty"`
	// 绑定关系编号

	BindId *string `json:"bind_id,omitempty"`
	// 绑定时间

	BindTime *sdktime.SdkTime `json:"bind_time,omitempty"`
	// API发布记录编号

	PublishId *string `json:"publish_id,omitempty"`
}

func (o AclBindApiInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AclBindApiInfo struct{}"
	}

	return strings.Join([]string{"AclBindApiInfo", string(data)}, " ")
}
