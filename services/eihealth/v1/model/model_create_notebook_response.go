package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNotebookResponse Response Object
type CreateNotebookResponse struct {

	// notebook ID
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateNotebookResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNotebookResponse struct{}"
	}

	return strings.Join([]string{"CreateNotebookResponse", string(data)}, " ")
}
