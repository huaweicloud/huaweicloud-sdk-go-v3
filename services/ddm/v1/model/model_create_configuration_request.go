package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateConfigurationRequest struct {

	// 名称。
	Name *string `json:"name,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 参数值。
	Values *interface{} `json:"values,omitempty"`
}

func (o CreateConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConfigurationRequest struct{}"
	}

	return strings.Join([]string{"CreateConfigurationRequest", string(data)}, " ")
}
