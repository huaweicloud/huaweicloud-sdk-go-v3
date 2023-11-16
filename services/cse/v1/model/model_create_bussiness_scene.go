package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBussinessScene 流量定义
type CreateBussinessScene struct {

	// 流量名称
	Name *string `json:"name,omitempty"`

	// 启用状态，支持enabled和disabled
	Status *string `json:"status,omitempty"`

	Selector *GovSelector `json:"selector,omitempty"`

	Spec *CreateBussinessSceneSpec `json:"spec,omitempty"`
}

func (o CreateBussinessScene) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBussinessScene struct{}"
	}

	return strings.Join([]string{"CreateBussinessScene", string(data)}, " ")
}
