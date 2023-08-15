package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSqlTemplatesRequest Request Object
type UpdateSqlTemplatesRequest struct {
	SqlId string `json:"sql_id"`

	Body *UpdateSqlTemplatesRequestBody `json:"body,omitempty"`
}

func (o UpdateSqlTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSqlTemplatesRequest struct{}"
	}

	return strings.Join([]string{"UpdateSqlTemplatesRequest", string(data)}, " ")
}
