package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateEventSchemaRequest struct {
	Body *CustomizeSchemaCreateReq `json:"body,omitempty"`
}

func (o CreateEventSchemaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEventSchemaRequest struct{}"
	}

	return strings.Join([]string{"CreateEventSchemaRequest", string(data)}, " ")
}
