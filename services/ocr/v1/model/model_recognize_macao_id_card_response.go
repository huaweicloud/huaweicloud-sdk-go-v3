package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeMacaoIdCardResponse Response Object
type RecognizeMacaoIdCardResponse struct {
	Result *MacaoIdCardResult `json:"result,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RecognizeMacaoIdCardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeMacaoIdCardResponse struct{}"
	}

	return strings.Join([]string{"RecognizeMacaoIdCardResponse", string(data)}, " ")
}
