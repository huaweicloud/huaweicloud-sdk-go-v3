package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateResourceConfigurationsReqBody struct {

	// 提权访问中的资源配置。
	ResourceConfigurations []ResourceConfiguration `json:"resource_configurations"`
}

func (o CreateResourceConfigurationsReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceConfigurationsReqBody struct{}"
	}

	return strings.Join([]string{"CreateResourceConfigurationsReqBody", string(data)}, " ")
}
