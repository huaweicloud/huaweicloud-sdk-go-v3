package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowRetrievalRequest struct {
	Body *QueryZoneReq `json:"body,omitempty"`
}

func (o ShowRetrievalRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRetrievalRequest struct{}"
	}

	return strings.Join([]string{"ShowRetrievalRequest", string(data)}, " ")
}
