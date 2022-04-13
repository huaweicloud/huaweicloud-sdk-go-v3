package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowRegisterServiceResponse struct {
	// 注册服务信息

	Services       *[]ServicesInfo `json:"services,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowRegisterServiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRegisterServiceResponse struct{}"
	}

	return strings.Join([]string{"ShowRegisterServiceResponse", string(data)}, " ")
}
