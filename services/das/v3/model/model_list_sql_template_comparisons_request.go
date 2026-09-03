package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSqlTemplateComparisonsRequest Request Object
type ListSqlTemplateComparisonsRequest struct {
	Body *ListSqlTemplateComparisonsRequestBody `json:"body,omitempty"`
}

func (o ListSqlTemplateComparisonsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlTemplateComparisonsRequest struct{}"
	}

	return strings.Join([]string{"ListSqlTemplateComparisonsRequest", string(data)}, " ")
}
