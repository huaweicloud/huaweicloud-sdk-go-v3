package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeHealthCodeResponse Response Object
type RecognizeHealthCodeResponse struct {
	Result *HealthCodeResult `json:"result,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RecognizeHealthCodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeHealthCodeResponse struct{}"
	}

	return strings.Join([]string{"RecognizeHealthCodeResponse", string(data)}, " ")
}
