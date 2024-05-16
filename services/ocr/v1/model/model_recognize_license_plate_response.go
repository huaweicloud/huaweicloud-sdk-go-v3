package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeLicensePlateResponse Response Object
type RecognizeLicensePlateResponse struct {

	// 调用成功时表示调用结果。  调用失败时无此字段。
	Result *[]LicensePlateResult `json:"result,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RecognizeLicensePlateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeLicensePlateResponse struct{}"
	}

	return strings.Join([]string{"RecognizeLicensePlateResponse", string(data)}, " ")
}
