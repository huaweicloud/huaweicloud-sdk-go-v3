package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApiConfig struct {

	// 资产名称
	Name *string `json:"name,omitempty"`
}

func (o ApiConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiConfig struct{}"
	}

	return strings.Join([]string{"ApiConfig", string(data)}, " ")
}
