package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserNotebookResponse Response Object
type ListUserNotebookResponse struct {

	// notebook总数
	Count *int32 `json:"count,omitempty"`

	// notebook详情列表
	Notebooks      *[]NotebookEntity `json:"notebooks,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListUserNotebookResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserNotebookResponse struct{}"
	}

	return strings.Join([]string{"ListUserNotebookResponse", string(data)}, " ")
}
