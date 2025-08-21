package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDirResponse Response Object
type CreateDirResponse struct {

	// 提交ID
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDirResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDirResponse struct{}"
	}

	return strings.Join([]string{"CreateDirResponse", string(data)}, " ")
}
