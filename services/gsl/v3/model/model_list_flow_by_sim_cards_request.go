package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListFlowBySimCardsRequest struct {
	Body *ListFlowBySimCardsReq `json:"body,omitempty"`
}

func (o ListFlowBySimCardsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlowBySimCardsRequest struct{}"
	}

	return strings.Join([]string{"ListFlowBySimCardsRequest", string(data)}, " ")
}
