package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGovPolicy 创建治理策略列表响应结构体
type CreateGovPolicy struct {

	// 治理策略名称
	Name *string `json:"name,omitempty"`

	Selector *GovSelector `json:"selector,omitempty"`

	// 治理策略定义内容
	Spec *interface{} `json:"spec,omitempty"`
}

func (o CreateGovPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGovPolicy struct{}"
	}

	return strings.Join([]string{"CreateGovPolicy", string(data)}, " ")
}
