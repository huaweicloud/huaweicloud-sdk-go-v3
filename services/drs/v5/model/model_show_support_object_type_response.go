package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSupportObjectTypeResponse Response Object
type ShowSupportObjectTypeResponse struct {

	// 全量任务是否支持对象选择。
	IsFullTransSupportObject *bool `json:"is_full_trans_support_object,omitempty"`

	// 增量任务是否支持对象选择。
	IsIncreTransSupportObject *bool `json:"is_incre_trans_support_object,omitempty"`

	// 全量加增量任务是否支持对象选择。
	IsFullIncreTransSupportObject *bool `json:"is_full_incre_trans_support_object,omitempty"`

	// 支持对象导入的引引擎。
	SupportObjectImportEngine *[]string `json:"support_object_import_engine,omitempty"`

	// 是否支持列映射。
	IsSupportColumnMapping *bool `json:"is_support_column_mapping,omitempty"`

	// 库是否支持搜索。
	IsDatabaseSupportSearch *bool `json:"is_database_support_search,omitempty"`

	// schema是否支持搜索。
	IsSchemaSupportSearch *bool `json:"is_schema_support_search,omitempty"`

	// 表是否支持搜索。
	IsTableSupportSearch *bool `json:"is_table_support_search,omitempty"`

	// 对象导入支持的文件大小。
	FileSize *string `json:"file_size,omitempty"`

	// 上一次选择迁移对象或者同步对象的方式。 - srcImportObject：当前任务上次选择的方式为导入方式
	PreviousSelect *string `json:"previous_select,omitempty"`

	// 对象导入类型。 - table：表级 - database：库级
	ImportLevel *string `json:"import_level,omitempty"`

	// 取值： - true： 当前任务上次选择列加工方式为导入方式 - false 或者 空：当前任务上次选择列加工方式为手动选择方式
	IsImportCloumn *bool `json:"is_import_cloumn,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowSupportObjectTypeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSupportObjectTypeResponse struct{}"
	}

	return strings.Join([]string{"ShowSupportObjectTypeResponse", string(data)}, " ")
}
