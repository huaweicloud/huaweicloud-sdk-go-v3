package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFlinkSqlJobTemplatesRequest Request Object
type ListFlinkSqlJobTemplatesRequest struct {

	// 返回的数据条数。默认为10。
	Limit *int32 `json:"limit,omitempty"`

	// 模板名字
	Name *string `json:"name,omitempty"`

	// 作业偏移量。
	Offset *int64 `json:"offset,omitempty"`

	// 查询结果排序，升序asc和降序desc两种可选，默认降序。
	Order *string `json:"order,omitempty"`

	Tags *string `json:"tags,omitempty"`
}

func (o ListFlinkSqlJobTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlinkSqlJobTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListFlinkSqlJobTemplatesRequest", string(data)}, " ")
}
