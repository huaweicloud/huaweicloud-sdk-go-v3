package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeExitEntryPermitRequest struct {
	Body *ExitEntryPermitRequestBody `json:"body,omitempty"`
}

func (o RecognizeExitEntryPermitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeExitEntryPermitRequest struct{}"
	}

	return strings.Join([]string{"RecognizeExitEntryPermitRequest", string(data)}, " ")
}
