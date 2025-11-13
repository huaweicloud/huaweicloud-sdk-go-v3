package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLivePlatformAccessTypeResponse Response Object
type ShowLivePlatformAccessTypeResponse struct {

	// 平台对接类型列表
	AccessTypes *[]AccessTypeEnum `json:"access_types,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowLivePlatformAccessTypeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLivePlatformAccessTypeResponse struct{}"
	}

	return strings.Join([]string{"ShowLivePlatformAccessTypeResponse", string(data)}, " ")
}
