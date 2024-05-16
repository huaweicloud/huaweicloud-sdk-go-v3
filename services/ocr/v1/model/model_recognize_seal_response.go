package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeSealResponse Response Object
type RecognizeSealResponse struct {
	Result *SealResult `json:"result,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RecognizeSealResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeSealResponse struct{}"
	}

	return strings.Join([]string{"RecognizeSealResponse", string(data)}, " ")
}
