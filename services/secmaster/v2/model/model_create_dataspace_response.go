package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDataspaceResponse Response Object
type CreateDataspaceResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDataspaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataspaceResponse struct{}"
	}

	return strings.Join([]string{"CreateDataspaceResponse", string(data)}, " ")
}
