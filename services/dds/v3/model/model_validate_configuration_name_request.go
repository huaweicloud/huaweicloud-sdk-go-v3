package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateConfigurationNameRequest Request Object
type ValidateConfigurationNameRequest struct {

	// 参数模板名称。
	Name string `json:"name"`
}

func (o ValidateConfigurationNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateConfigurationNameRequest struct{}"
	}

	return strings.Join([]string{"ValidateConfigurationNameRequest", string(data)}, " ")
}
