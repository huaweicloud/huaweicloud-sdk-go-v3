package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateKieReq struct {

	// 配置项的id。
	Id *string `json:"id,omitempty"`

	// 配置项的key。
	Key *string `json:"key,omitempty"`

	// 配置项的标签
	Labels *interface{} `json:"labels,omitempty"`

	// 配置项的值。
	Value *string `json:"value,omitempty"`

	// 配置项value的类型。
	ValueType *string `json:"value_type,omitempty"`

	// 配置项的状态。
	Status *string `json:"status,omitempty"`
}

func (o CreateKieReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKieReq struct{}"
	}

	return strings.Join([]string{"CreateKieReq", string(data)}, " ")
}
