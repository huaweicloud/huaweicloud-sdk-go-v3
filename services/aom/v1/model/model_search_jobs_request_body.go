package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 分页查询作业列表请求参数
type SearchJobsRequestBody struct {

	// name为作业名称。
	Name *string `json:"name,omitempty"`

	// 当前页，查询的当前页，page_num为正整数，不能是0和负数，当输入参数为负数，0和大于1000，自动修正参数为1，默认值是1（用户不传，值是1）。
	PageNum *int32 `json:"page_num,omitempty"`

	// 每页显示的条数，每页查询的总条数，page_size为正整数，不能是0和负数，当输入参数为负数，0和大于101，自动修正参数为10，默认值是10（用户不传时，值是10）。
	PageSize *int32 `json:"page_size,omitempty"`

	// 需要排序的字段(默认为更新时间),支持字段有name,create_time,update_time。
	OrderByColumn string `json:"order_by_column"`

	// 排序规则(默认降序) 传入升序或降序，升序：ASC，降序：DESC。
	SortOrder string `json:"sort_order"`

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o SearchJobsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchJobsRequestBody struct{}"
	}

	return strings.Join([]string{"SearchJobsRequestBody", string(data)}, " ")
}
