package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSqlTemplatesResponse Response Object
type ListSqlTemplatesResponse struct {

	// SQL模板列表
	TplList        *[]Tpl `json:"tpl_list,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSqlTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListSqlTemplatesResponse", string(data)}, " ")
}
