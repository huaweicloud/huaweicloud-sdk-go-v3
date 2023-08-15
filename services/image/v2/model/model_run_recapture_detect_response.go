package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunRecaptureDetectResponse Response Object
type RunRecaptureDetectResponse struct {
	Result         *RecaptureDetectResponseResult `json:"result,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o RunRecaptureDetectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunRecaptureDetectResponse struct{}"
	}

	return strings.Join([]string{"RunRecaptureDetectResponse", string(data)}, " ")
}
