package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteSqlJobTemplatesRequest Request Object
type BatchDeleteSqlJobTemplatesRequest struct {
	Body *BatchDeleteSqlJobTemplatesRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteSqlJobTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSqlJobTemplatesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteSqlJobTemplatesRequest", string(data)}, " ")
}
