package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTablesResponse Response Object
type ListTablesResponse struct {

	// 表的信息。
	Tables *[]Table `json:"tables,omitempty"`

	// 数据表总数。
	Count          *int64 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTablesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTablesResponse struct{}"
	}

	return strings.Join([]string{"ListTablesResponse", string(data)}, " ")
}
