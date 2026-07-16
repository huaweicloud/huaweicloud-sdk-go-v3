package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNotebookTagsResponse Response Object
type CreateNotebookTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateNotebookTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNotebookTagsResponse struct{}"
	}

	return strings.Join([]string{"CreateNotebookTagsResponse", string(data)}, " ")
}
