package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowApplicationConfigurationResponse Response Object
type ShowApplicationConfigurationResponse struct {

	// 应用配置列表。
	Configuration  *[]ApplicationListConfigConfiguration1 `json:"configuration,omitempty"`
	HttpStatusCode int                                    `json:"-"`
}

func (o ShowApplicationConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApplicationConfigurationResponse struct{}"
	}

	return strings.Join([]string{"ShowApplicationConfigurationResponse", string(data)}, " ")
}
