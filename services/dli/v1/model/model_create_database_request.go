package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDatabaseRequest Request Object
type CreateDatabaseRequest struct {
	Body *CreateDatabaseRequestBody `json:"body,omitempty"`
}

func (o CreateDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatabaseRequest struct{}"
	}

	return strings.Join([]string{"CreateDatabaseRequest", string(data)}, " ")
}
