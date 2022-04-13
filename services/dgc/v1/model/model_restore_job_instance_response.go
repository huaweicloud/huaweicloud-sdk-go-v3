package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RestoreJobInstanceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RestoreJobInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreJobInstanceResponse struct{}"
	}

	return strings.Join([]string{"RestoreJobInstanceResponse", string(data)}, " ")
}
