package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateProxyNewConfigRequestBody struct {

	// 修改的配置信息。
	Configurations []UpdateProxyConfigurationItem `json:"configurations"`
}

func (o UpdateProxyNewConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProxyNewConfigRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateProxyNewConfigRequestBody", string(data)}, " ")
}
