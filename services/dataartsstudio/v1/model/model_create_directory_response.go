package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDirectoryResponse Response Object
type CreateDirectoryResponse struct {
	Data           *CreateDirectoryResultData `json:"data,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o CreateDirectoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDirectoryResponse struct{}"
	}

	return strings.Join([]string{"CreateDirectoryResponse", string(data)}, " ")
}
