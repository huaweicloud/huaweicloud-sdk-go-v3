package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSqlTemplatesRequest Request Object
type ListSqlTemplatesRequest struct {
	Body *ListSqlTemplatesRequestBody `json:"body,omitempty"`
}

func (o ListSqlTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListSqlTemplatesRequest", string(data)}, " ")
}
