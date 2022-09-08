package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RecognizeExitEntryPermitResponse struct {
	Result         *ExitEntryPermitResult `json:"result,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o RecognizeExitEntryPermitResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeExitEntryPermitResponse struct{}"
	}

	return strings.Join([]string{"RecognizeExitEntryPermitResponse", string(data)}, " ")
}
