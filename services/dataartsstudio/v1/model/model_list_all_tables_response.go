package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllTablesResponse Response Object
type ListAllTablesResponse struct {
	Data           *ListAllTablesResultData `json:"data,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListAllTablesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllTablesResponse struct{}"
	}

	return strings.Join([]string{"ListAllTablesResponse", string(data)}, " ")
}
