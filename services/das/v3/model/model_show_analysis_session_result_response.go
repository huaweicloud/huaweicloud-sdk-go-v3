package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAnalysisSessionResultResponse Response Object
type ShowAnalysisSessionResultResponse struct {
	Body           *[]ShowAnalysisSessionResultResp `json:"body,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o ShowAnalysisSessionResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAnalysisSessionResultResponse struct{}"
	}

	return strings.Join([]string{"ShowAnalysisSessionResultResponse", string(data)}, " ")
}
