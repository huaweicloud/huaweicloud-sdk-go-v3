package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeMacaoIdCardRequest struct {
	Body *MacaoIdCardRequestBody `json:"body,omitempty"`
}

func (o RecognizeMacaoIdCardRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeMacaoIdCardRequest struct{}"
	}

	return strings.Join([]string{"RecognizeMacaoIdCardRequest", string(data)}, " ")
}
