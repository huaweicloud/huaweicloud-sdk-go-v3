package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSqlTemplatesRequest Request Object
type CreateSqlTemplatesRequest struct {
	Body *CreateSqlTemplatesRequestBody `json:"body,omitempty"`
}

func (o CreateSqlTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSqlTemplatesRequest struct{}"
	}

	return strings.Join([]string{"CreateSqlTemplatesRequest", string(data)}, " ")
}
