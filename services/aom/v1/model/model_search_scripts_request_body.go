package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SearchScriptsRequestBody struct {

	// 查询接收的的参数，版本管理时，name为脚本名称（版本管理查询时name不能为空），页面查询时，name为接收模糊查询的参数，name是null，表示查询所有默认脚本。
	Name *string `json:"name,omitempty"`

	// 查询规则，default表示脚本管理主页查询，包括模糊查询，no_default就是版本管理,当defalult不输入（参数为空）进行查询时，默认是页面查询,废弃字段，传入不影响使用。
	IsDefault *string `json:"is_default,omitempty"`

	// 创建人，默认按照创建人搜说脚本。
	CreateBy *string `json:"create_by,omitempty"`

	// 版本管理时需要查询的脚本id。
	ScriptId *string `json:"script_id,omitempty"`

	// 当前页，查询的当前页，page_num为正整数，不能是0和负数，当输入参数为负数，0和大于1000，自动修正参数为1，默认值是1（用户不传，值是1）。
	PageNum *int32 `json:"page_num,omitempty"`

	// 每页显示的条数，每页查询的总条数，page_size为正整数，不能是0和负数，当输入参数为负数，0和大于101，自动修正参数为10，默认值是10（用户不传时，值是10）。
	PageSize *int32 `json:"page_size,omitempty"`

	// 项目id，版本管理参数，当时page_total为total,返回出改脚本名称下所有的脚本对象,传入其他参数无意义，，当前进行 版本管理时，需要分页显示脚本名称下所有对象，传入空值即可。
	ProjectId string `json:"project_id"`

	// 需要排序的字段(默认为更新时间),支持字段有name,create_time和update_time。
	OrderByColumn *string `json:"order_by_column,omitempty"`

	// 排序规则(默认降序) 传入升序或降序，升序：ASC，降序：DESC。
	SortOrder *string `json:"sort_order,omitempty"`

	// 版本管理参数，当时page_total为total,返回出改脚本名称下所有的脚本对象,传入其他参数无意义，，当前进行 版本管理时，需要分页显示脚本名称下所有对象，传入空值即可。
	PageTotal *string `json:"page_total,omitempty"`

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
