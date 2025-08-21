package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFileUpperTreeEntriesResponse Response Object
type ListFileUpperTreeEntriesResponse struct {

	// 获取当前文件上级树结构
	Body           *[]TreeDto `json:"body,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListFileUpperTreeEntriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFileUpperTreeEntriesResponse struct{}"
	}

	return strings.Join([]string{"ListFileUpperTreeEntriesResponse", string(data)}, " ")
}
