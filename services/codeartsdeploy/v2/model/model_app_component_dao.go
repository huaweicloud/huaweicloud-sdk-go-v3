package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AppComponentDao 应用和AOM应用组件对应关系
type AppComponentDao struct {

	// 部署任务id
	TaskId *string `json:"task_id,omitempty"`

	// AOM应用id
	AppId *string `json:"app_id,omitempty"`

	// AOM应用名称
	AppName *string `json:"app_name,omitempty"`

	// AOM应用组件id
	CompId *string `json:"comp_id,omitempty"`

	// AOM应用组件名称
	CompName *string `json:"comp_name,omitempty"`

	// 租户id
	DomainId *string `json:"domain_id,omitempty"`

	// 局点信息
	Region *string `json:"region,omitempty"`

	// AOM应用组件是否生效，0表示初始化，1表示执行成功，已生效
	State *string `json:"state,omitempty"`
}

func (o AppComponentDao) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppComponentDao struct{}"
	}

	return strings.Join([]string{"AppComponentDao", string(data)}, " ")
}
