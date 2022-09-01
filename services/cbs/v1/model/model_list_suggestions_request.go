package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListSuggestionsRequest struct {

	// qabot编号，UUID格式。
	QabotId string `json:"qabot_id" xml:"qabot_id"`

	Body *PostSuggestionsReq `json:"body,omitempty" xml:"body"`
}

func (o ListSuggestionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSuggestionsRequest struct{}"
	}

	return strings.Join([]string{"ListSuggestionsRequest", string(data)}, " ")
}
