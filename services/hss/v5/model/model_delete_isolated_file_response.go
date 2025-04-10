package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteIsolatedFileResponse Response Object
type DeleteIsolatedFileResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteIsolatedFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIsolatedFileResponse struct{}"
	}

	return strings.Join([]string{"DeleteIsolatedFileResponse", string(data)}, " ")
}
