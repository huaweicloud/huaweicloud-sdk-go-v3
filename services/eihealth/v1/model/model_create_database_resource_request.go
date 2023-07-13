package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDatabaseResourceRequest Request Object
type CreateDatabaseResourceRequest struct {
	Body *CreateDatabaseResourceReq `json:"body,omitempty"`
}

func (o CreateDatabaseResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatabaseResourceRequest struct{}"
	}

	return strings.Join([]string{"CreateDatabaseResourceRequest", string(data)}, " ")
}
