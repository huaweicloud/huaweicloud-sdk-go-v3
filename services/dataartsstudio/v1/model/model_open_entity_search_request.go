package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OpenEntitySearchRequest 资产查询条件
type OpenEntitySearchRequest struct {

	// 查询关键字
	Query string `json:"query"`

	// 分类名称 List<String>
	TypeNames []string `json:"type_names"`

	// List<String> 连接名称
	ConnectionNames *[]string `json:"connection_names,omitempty"`

	// 查询关键字是否匹配资产的名称描述信息，true:匹配所有属性，false:只匹配名称、描述，默认false
	SearchAllAttributes *bool `json:"search_all_attributes,omitempty"`

	// List<String> 标签的名称
	Tags *[]string `json:"tags,omitempty"`

	// 分页显示每页返回结果数。默认值，10
	Limit int32 `json:"limit"`

	// 偏移量，默认值，0
	Offset *int32 `json:"offset,omitempty"`

	// key当前支持Table，value可为以下中的一个或多个：rowCounts、tableSize、database、schema、namespace、ddlUpdateTime、dataUpdateTime、ddlCreateTime Map<String,Set<String>>
	Attributes *interface{} `json:"attributes,omitempty"`

	FilterCriteria *FilterCriteria `json:"filter_criteria,omitempty"`

	TimeRange *TimeRange `json:"time_range,omitempty"`

	// scroll_id
	ScrollId *string `json:"scroll_id,omitempty"`

	// List<String> 安全级别
	SecurityLevels *[]string `json:"security_levels,omitempty"`

	// 是否导入
	IsImport *bool `json:"is_import,omitempty"`

	// List<String> 分类
	Classifications *[]string `json:"classifications,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`
}

func (o OpenEntitySearchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenEntitySearchRequest struct{}"
	}

	return strings.Join([]string{"OpenEntitySearchRequest", string(data)}, " ")
}
