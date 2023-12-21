package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEcsQuotaResponse Response Object
type ShowEcsQuotaResponse struct {

	// 支持IPv6云堡垒机实例规格资源状态。 - sellout：售罄 - normal：正常商用
	StatusV6 *string `json:"status_v6,omitempty"`

	// 云堡垒机实例规格资源状态。 - sellout：售罄 - normal：正常商用
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowEcsQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEcsQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowEcsQuotaResponse", string(data)}, " ")
}
