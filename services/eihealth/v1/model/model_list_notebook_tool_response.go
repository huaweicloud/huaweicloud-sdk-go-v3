package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNotebookToolResponse Response Object
type ListNotebookToolResponse struct {

	// 总个数
	Count *int32 `json:"count,omitempty"`

	// tool详情列表
	Tools          *[]NotebookToolDto `json:"tools,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListNotebookToolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotebookToolResponse struct{}"
	}

	return strings.Join([]string{"ListNotebookToolResponse", string(data)}, " ")
}
