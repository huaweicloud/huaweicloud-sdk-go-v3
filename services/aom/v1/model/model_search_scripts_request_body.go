package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SearchScriptsRequestBody struct {

	// 查询接收的参数，版本管理时，name为脚本名称（版本管理查询时name不能为空），脚本管理页面查询时，name为接收模糊查询的参数，name是null，表示查询所有默认脚本。
	Name *string `json:"name,omitempty"`

	// 查询规则，如果是类型为default，则为模糊查询和脚本管理主页展示，no_default为版本管理。
	IsDefault *string `json:"is_default,omitempty"`

	// 创建人，默认按照创建人搜索脚本。
	CreateBy *string `json:"create_by,omitempty"`

	// 版本管理时需要查询的脚本id。
	ScriptId *string `json:"script_id,omitempty"`

	// page_num为正整数。
	PageNum *int32 `json:"page_num,omitempty"`

	// 每页显示的条数，默认值是10。
	PageSize *int32 `json:"page_size,omitempty"`

	// 项目id。
	ProjectId *string `json:"project_id,omitempty"`

	// 需要排序的字段(默认为更新时间),支持字段有name,create_time和update_time。
	OrderByColumn string `json:"order_by_column"`

	// 排序规则(默认降序) 传入升序或降序，升序：ASC，降序：DESC。
	SortOrder *string `json:"sort_order,omitempty"`

	// 企业项目id，根据企业项目id搜索。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o SearchScriptsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchScriptsRequestBody struct{}"
	}

	return strings.Join([]string{"SearchScriptsRequestBody", string(data)}, " ")
}
