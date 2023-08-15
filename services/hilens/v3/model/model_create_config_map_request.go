package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConfigMapRequest Request Object
type CreateConfigMapRequest struct {

	// 服务名称，hilens或者ief，默认hilens
	Provider *string `json:"provider,omitempty"`

	Body *ConfigMapModelBoxDto `json:"body,omitempty"`
}

func (o CreateConfigMapRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConfigMapRequest struct{}"
	}

	return strings.Join([]string{"CreateConfigMapRequest", string(data)}, " ")
}
