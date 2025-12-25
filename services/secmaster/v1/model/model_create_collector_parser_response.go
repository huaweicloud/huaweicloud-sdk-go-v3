package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCollectorParserResponse Response Object
type CreateCollectorParserResponse struct {

	// UUID
	ParserId       *string `json:"parser_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateCollectorParserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCollectorParserResponse struct{}"
	}

	return strings.Join([]string{"CreateCollectorParserResponse", string(data)}, " ")
}
