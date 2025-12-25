package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTablesResponse Response Object
type ListTablesResponse struct {

	// 表数量
	Count *int64 `json:"count,omitempty"`

	// 表具体内容
	Records        *[]IsapTable `json:"records,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListTablesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTablesResponse struct{}"
	}

	return strings.Join([]string{"ListTablesResponse", string(data)}, " ")
}
