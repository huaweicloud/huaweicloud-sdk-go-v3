package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchErrorInfoSource4ApiResponse Response Object
type SearchErrorInfoSource4ApiResponse struct {

	// 错误信息条件列表
	EntryNames     *[]string `json:"entry_names,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o SearchErrorInfoSource4ApiResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchErrorInfoSource4ApiResponse struct{}"
	}

	return strings.Join([]string{"SearchErrorInfoSource4ApiResponse", string(data)}, " ")
}
