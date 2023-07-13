package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFileResponse Response Object
type DeleteFileResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFileResponse struct{}"
	}

	return strings.Join([]string{"DeleteFileResponse", string(data)}, " ")
}
