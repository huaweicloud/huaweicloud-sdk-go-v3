package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListFlowBySimCardsResponse struct {
	Body           *[]SimCardsFlowVo `json:"body,omitempty" xml:"body"`
	HttpStatusCode int               `json:"-"`
}

func (o ListFlowBySimCardsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlowBySimCardsResponse struct{}"
	}

	return strings.Join([]string{"ListFlowBySimCardsResponse", string(data)}, " ")
}
