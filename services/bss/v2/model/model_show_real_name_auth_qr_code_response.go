package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRealNameAuthQrCodeResponse Response Object
type ShowRealNameAuthQrCodeResponse struct {

	// 人脸实名认证二维码地址。该二维码仅限单次使用，扫描后将自动失效。若未在10分钟内完成扫描，系统将自动作废。
	QrCodeUrl      *string `json:"qr_code_url,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowRealNameAuthQrCodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRealNameAuthQrCodeResponse struct{}"
	}

	return strings.Join([]string{"ShowRealNameAuthQrCodeResponse", string(data)}, " ")
}
