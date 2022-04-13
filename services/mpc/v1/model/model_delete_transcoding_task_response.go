package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteTranscodingTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTranscodingTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTranscodingTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteTranscodingTaskResponse", string(data)}, " ")
}
