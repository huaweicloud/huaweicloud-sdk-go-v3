package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 部署任务和应用组件对应关系
type AppComponentDao struct {

	// 部署任务id
	TaskId *string `json:"task_id,omitempty" xml:"task_id"`

	// 应用id
	AppId *string `json:"app_id,omitempty" xml:"app_id"`

	// 应用名称
	AppName *string `json:"app_name,omitempty" xml:"app_name"`

	// 组件id
	CompId *string `json:"comp_id,omitempty" xml:"comp_id"`

	// 组件名称
	CompName *string `json:"comp_name,omitempty" xml:"comp_name"`

	// 租户ID
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id"`

	// 局点信息
	Region *string `json:"region,omitempty" xml:"region"`

	// 组件是否生效，0表示初始化，1表示执行成功，已生效
	State *string `json:"state,omitempty" xml:"state"`
}

func (o AppComponentDao) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppComponentDao struct{}"
	}

	return strings.Join([]string{"AppComponentDao", string(data)}, " ")
}
