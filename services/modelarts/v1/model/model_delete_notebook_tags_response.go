package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNotebookTagsResponse Response Object
type DeleteNotebookTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteNotebookTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNotebookTagsResponse struct{}"
	}

	return strings.Join([]string{"DeleteNotebookTagsResponse", string(data)}, " ")
}
