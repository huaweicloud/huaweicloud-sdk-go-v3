package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopOrStartNotebookResponse Response Object
type StopOrStartNotebookResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopOrStartNotebookResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopOrStartNotebookResponse struct{}"
	}

	return strings.Join([]string{"StopOrStartNotebookResponse", string(data)}, " ")
}
