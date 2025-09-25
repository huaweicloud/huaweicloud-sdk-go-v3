package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSqlTraceResponse Response Object
type ListSqlTraceResponse struct {
	Body           *[]NodeExecutionInfoResult `json:"body,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListSqlTraceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlTraceResponse struct{}"
	}

	return strings.Join([]string{"ListSqlTraceResponse", string(data)}, " ")
}
