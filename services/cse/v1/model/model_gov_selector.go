package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GovSelector 治理策略下发范围
type GovSelector struct {

	// 所属环境
	Environment *string `json:"environment,omitempty"`

	// 所属应用
	App *string `json:"app,omitempty"`

	// 可选，治理下发到微服务级别
	Service *string `json:"service,omitempty"`
}

func (o GovSelector) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GovSelector struct{}"
	}

	return strings.Join([]string{"GovSelector", string(data)}, " ")
}
