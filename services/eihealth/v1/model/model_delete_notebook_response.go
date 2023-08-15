package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNotebookResponse Response Object
type DeleteNotebookResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteNotebookResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNotebookResponse struct{}"
	}

	return strings.Join([]string{"DeleteNotebookResponse", string(data)}, " ")
}
