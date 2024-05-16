package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeVietnamIdCardResponse Response Object
type RecognizeVietnamIdCardResponse struct {
	Result *VietnamIdCardResult `json:"result,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RecognizeVietnamIdCardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeVietnamIdCardResponse struct{}"
	}

	return strings.Join([]string{"RecognizeVietnamIdCardResponse", string(data)}, " ")
}
