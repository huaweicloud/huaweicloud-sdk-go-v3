package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNotebookRequest Request Object
type CreateNotebookRequest struct {
	Body *NotebookCreateRequest `json:"body,omitempty"`
}

func (o CreateNotebookRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNotebookRequest struct{}"
	}

	return strings.Join([]string{"CreateNotebookRequest", string(data)}, " ")
}
