package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeThailandLicensePlateResponse Response Object
type RecognizeThailandLicensePlateResponse struct {

	// 调用成功时表示调用结果。 调用失败时无此字段
	Result *[]ThailandLicensePlateItem `json:"result,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RecognizeThailandLicensePlateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeThailandLicensePlateResponse struct{}"
	}

	return strings.Join([]string{"RecognizeThailandLicensePlateResponse", string(data)}, " ")
}
