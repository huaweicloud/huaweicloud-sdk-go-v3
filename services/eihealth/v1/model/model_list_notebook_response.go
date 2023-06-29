package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNotebookResponse Response Object
type ListNotebookResponse struct {

	// notebook总数
	Count *int32 `json:"count,omitempty"`

	// notebook详情列表
	Notebooks      *[]NotebookEntity `json:"notebooks,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListNotebookResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotebookResponse struct{}"
	}

	return strings.Join([]string{"ListNotebookResponse", string(data)}, " ")
}
