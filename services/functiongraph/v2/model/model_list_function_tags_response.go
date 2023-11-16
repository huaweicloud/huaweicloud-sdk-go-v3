package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFunctionTagsResponse Response Object
type ListFunctionTagsResponse struct {

	// 标签列表
	Tags *[]KvItem `json:"tags,omitempty"`

	// 系统标签列表
	SysTags        *[]KvItem `json:"sys_tags,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListFunctionTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFunctionTagsResponse struct{}"
	}

	return strings.Join([]string{"ListFunctionTagsResponse", string(data)}, " ")
}
