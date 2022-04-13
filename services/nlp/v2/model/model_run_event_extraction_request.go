package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunEventExtractionRequest struct {
	Body *PostEventExtractionReq `json:"body,omitempty"`
}

func (o RunEventExtractionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunEventExtractionRequest struct{}"
	}

	return strings.Join([]string{"RunEventExtractionRequest", string(data)}, " ")
}
