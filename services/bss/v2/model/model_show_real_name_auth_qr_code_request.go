package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRealNameAuthQrCodeRequest Request Object
type ShowRealNameAuthQrCodeRequest struct {
}

func (o ShowRealNameAuthQrCodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRealNameAuthQrCodeRequest struct{}"
	}

	return strings.Join([]string{"ShowRealNameAuthQrCodeRequest", string(data)}, " ")
}
