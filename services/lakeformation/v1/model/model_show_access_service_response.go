package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAccessServiceResponse Response Object
type ShowAccessServiceResponse struct {

	// 接入服务授权信息列表
	AccessServices *[]AccessServiceInfo `json:"access_services,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAccessServiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAccessServiceResponse struct{}"
	}

	return strings.Join([]string{"ShowAccessServiceResponse", string(data)}, " ")
}
