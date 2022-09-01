package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeCambodianIdCardRequest struct {
	Body *CambodianIdCardRequestBody `json:"body,omitempty" xml:"body"`
}

func (o RecognizeCambodianIdCardRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeCambodianIdCardRequest struct{}"
	}

	return strings.Join([]string{"RecognizeCambodianIdCardRequest", string(data)}, " ")
}
