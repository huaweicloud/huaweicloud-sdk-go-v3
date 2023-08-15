package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFlowBySimCardsResponse Response Object
type ListFlowBySimCardsResponse struct {
	Body           *[]SimCardsFlowVo `json:"body,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListFlowBySimCardsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlowBySimCardsResponse struct{}"
	}

	return strings.Join([]string{"ListFlowBySimCardsResponse", string(data)}, " ")
}
