package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchCodeTablesResponse Response Object
type SearchCodeTablesResponse struct {
	Data           *SearchCodeTablesResultData `json:"data,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o SearchCodeTablesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchCodeTablesResponse struct{}"
	}

	return strings.Join([]string{"SearchCodeTablesResponse", string(data)}, " ")
}
