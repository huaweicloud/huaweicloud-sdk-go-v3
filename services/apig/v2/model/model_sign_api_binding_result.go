package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SignApiBindingResult struct {
	// API与签名密钥的绑定关系列表

	Bindings *[]SignApiBindingInfo `json:"bindings,omitempty"`
}

func (o SignApiBindingResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SignApiBindingResult struct{}"
	}

	return strings.Join([]string{"SignApiBindingResult", string(data)}, " ")
}
