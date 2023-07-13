package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNotebookResponse Response Object
type UpdateNotebookResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateNotebookResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNotebookResponse struct{}"
	}

	return strings.Join([]string{"UpdateNotebookResponse", string(data)}, " ")
}
