package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AppFilterDto struct {

	// 应用id
	AppId *string `json:"app_id,omitempty"`

	// 应用名称
	AppName *string `json:"app_name,omitempty"`

	// 应用版本
	AppVersion *string `json:"app_version,omitempty"`

	// 计算节点标签
	AppNodeLabels *[]string `json:"app_node_labels,omitempty"`
}

func (o AppFilterDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppFilterDto struct{}"
	}

	return strings.Join([]string{"AppFilterDto", string(data)}, " ")
}
