package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSqlTemplateComparisonsResponse Response Object
type ListSqlTemplateComparisonsResponse struct {

	// SQL模板对比列表
	SqlTplCmpDtoList *[]SqlTplCmp `json:"sql_tpl_cmp_dto_list,omitempty"`
	HttpStatusCode   int          `json:"-"`
}

func (o ListSqlTemplateComparisonsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlTemplateComparisonsResponse struct{}"
	}

	return strings.Join([]string{"ListSqlTemplateComparisonsResponse", string(data)}, " ")
}
