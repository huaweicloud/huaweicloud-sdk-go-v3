package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSqlTemplatesRequest Request Object
type ShowSqlTemplatesRequest struct {

	// 用于过滤SQL模板的名字。
	Keyword *string `json:"keyword,omitempty"`
}

func (o ShowSqlTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ShowSqlTemplatesRequest", string(data)}, " ")
}
