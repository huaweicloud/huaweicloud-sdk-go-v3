package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllVersionByVersionIdResponse Response Object
type ListAllVersionByVersionIdResponse struct {

	// 查询结果集合。
	Elements *[]ScriptVersion `json:"elements,omitempty"`

	// 查询到的结果数量。
	TotalElements  *int32 `json:"total_elements,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAllVersionByVersionIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllVersionByVersionIdResponse struct{}"
	}

	return strings.Join([]string{"ListAllVersionByVersionIdResponse", string(data)}, " ")
}
