package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDimensionLogicTablesResponse Response Object
type ListDimensionLogicTablesResponse struct {
	Data           *ListDimensionLogicTablesResultData `json:"data,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o ListDimensionLogicTablesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDimensionLogicTablesResponse struct{}"
	}

	return strings.Join([]string{"ListDimensionLogicTablesResponse", string(data)}, " ")
}
