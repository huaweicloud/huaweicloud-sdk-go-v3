package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDimensionLogicTablesResponse Response Object
type ListDimensionLogicTablesResponse struct {

	// 返回的数据信息。
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListDimensionLogicTablesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDimensionLogicTablesResponse struct{}"
	}

	return strings.Join([]string{"ListDimensionLogicTablesResponse", string(data)}, " ")
}
