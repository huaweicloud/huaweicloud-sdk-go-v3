package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFactLogicTablesResponse Response Object
type ListFactLogicTablesResponse struct {
	Data           *ListFactLogicTablesResultData `json:"data,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ListFactLogicTablesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFactLogicTablesResponse struct{}"
	}

	return strings.Join([]string{"ListFactLogicTablesResponse", string(data)}, " ")
}
