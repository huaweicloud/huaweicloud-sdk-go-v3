package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SupportedService 支持服务专属凭证的云服务信息。
type SupportedService struct {

	// 云服务名称。
	ServiceName string `json:"service_name"`

	// 关联的服务的展示名称（受语言控制）。
	DisplayName string `json:"display_name"`
}

func (o SupportedService) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupportedService struct{}"
	}

	return strings.Join([]string{"SupportedService", string(data)}, " ")
}
