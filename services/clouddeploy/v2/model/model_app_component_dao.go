package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 部署任务和应用组件对应关系
type AppComponentDao struct {
	// 部署任务id

	TaskId *string `json:"task_id,omitempty"`
	// 应用id

	AppId *string `json:"app_id,omitempty"`
	// 应用名称

	AppName *string `json:"app_name,omitempty"`
	// 组件id

	CompId *string `json:"comp_id,omitempty"`
	// 组件名称

	CompName *string `json:"comp_name,omitempty"`
	// 租户ID

	DomainId *string `json:"domain_id,omitempty"`
	// 局点信息

	Region *string `json:"region,omitempty"`
	// 组件是否生效，0表示初始化，1表示执行成功，已生效

	State *string `json:"state,omitempty"`
}

func (o AppComponentDao) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppComponentDao struct{}"
	}

	return strings.Join([]string{"AppComponentDao", string(data)}, " ")
}
