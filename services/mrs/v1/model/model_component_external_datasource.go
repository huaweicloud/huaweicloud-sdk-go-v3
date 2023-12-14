package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ComponentExternalDatasource 组件外部数据源
type ComponentExternalDatasource struct {

	// 外部数据源名称
	Name *string `json:"name,omitempty"`

	// 外部数据源类型
	Types *[]string `json:"types,omitempty"`
}

func (o ComponentExternalDatasource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentExternalDatasource struct{}"
	}

	return strings.Join([]string{"ComponentExternalDatasource", string(data)}, " ")
}
