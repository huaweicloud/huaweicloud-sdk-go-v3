package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskCheckParamters struct {
	// 检查工具ID

	CheckId *int32 `json:"check_id,omitempty"`
	// 编译参数名称

	Name *string `json:"name,omitempty"`
	// 检查参数配置信息

	CheckerConfigs *[]CheckConfigInfo `json:"checker_configs,omitempty"`
}

func (o TaskCheckParamters) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskCheckParamters struct{}"
	}

	return strings.Join([]string{"TaskCheckParamters", string(data)}, " ")
}
