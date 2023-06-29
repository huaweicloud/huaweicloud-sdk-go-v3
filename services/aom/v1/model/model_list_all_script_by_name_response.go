package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllScriptByNameResponse Response Object
type ListAllScriptByNameResponse struct {

	// 查询结果集合。
	Elements *[]Script `json:"elements,omitempty"`

	// 查询到的结果数量。
	TotalElements  *int32 `json:"total_elements,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAllScriptByNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllScriptByNameResponse struct{}"
	}

	return strings.Join([]string{"ListAllScriptByNameResponse", string(data)}, " ")
}
