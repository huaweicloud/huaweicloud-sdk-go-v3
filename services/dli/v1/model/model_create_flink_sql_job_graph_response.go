package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFlinkSqlJobGraphResponse Response Object
type CreateFlinkSqlJobGraphResponse struct {
	IsSuccess *bool `json:"is_success,omitempty"`

	Message *string `json:"message,omitempty"`

	ErrorCode *string `json:"error_code,omitempty"`

	// 静态流图的描述信息
	StreamGraph    *string `json:"stream_graph,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateFlinkSqlJobGraphResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlinkSqlJobGraphResponse struct{}"
	}

	return strings.Join([]string{"CreateFlinkSqlJobGraphResponse", string(data)}, " ")
}
