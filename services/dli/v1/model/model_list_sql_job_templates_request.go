package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSqlJobTemplatesRequest Request Object
type ListSqlJobTemplatesRequest struct {

	// 用于过滤SQL模板的名字。
	Keyword *string `json:"keyword,omitempty"`
}

func (o ListSqlJobTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlJobTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListSqlJobTemplatesRequest", string(data)}, " ")
}
