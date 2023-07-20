package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSqlTemplatesRequest Request Object
type DeleteSqlTemplatesRequest struct {
	Body *DeleteSqlTemplatesRequestBody `json:"body,omitempty"`
}

func (o DeleteSqlTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSqlTemplatesRequest struct{}"
	}

	return strings.Join([]string{"DeleteSqlTemplatesRequest", string(data)}, " ")
}
