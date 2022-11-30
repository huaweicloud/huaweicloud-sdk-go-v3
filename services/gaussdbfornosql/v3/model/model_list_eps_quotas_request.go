package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListEpsQuotasRequest struct {

	// 企业项目名称。支持模糊搜索，若不指定则返回所有企业项目配额。
	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`

	// 索引位置，偏移量。    - 从第一条数据偏移offset条数据后开始查询，默认为0。   - 取值必须为数字，且不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询个数上限值。  - 取值范围：1~100。 - 不传该参数时，默认查询前100条信息。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListEpsQuotasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEpsQuotasRequest struct{}"
	}

	return strings.Join([]string{"ListEpsQuotasRequest", string(data)}, " ")
}
