package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TableLineCompareResultInfo 表级别对比结果信息体
type TableLineCompareResultInfo struct {

	// 源库表名称
	SourceTableName *string `json:"source_table_name,omitempty"`

	// 源库表行数
	SourceRowNum *int64 `json:"source_row_num,omitempty"`

	// 目标库表名称
	TargetTableName *string `json:"target_table_name,omitempty"`

	// 目标库表行数
	TargetRowNum *int64 `json:"target_row_num,omitempty"`

	// 行数差异值
	DifferenceRowNum *int64 `json:"difference_row_num,omitempty"`

	// 对比状态。 - 0：对比不一致 - 2：对比一致 - 3：目标库表不存在 - 4：对比失败 - 5：正在对比中 - 6：等待对比中 - 7：任务已取消 - 8：源库为空 - 9：目标库为空 - 10：源库和目标库都为空 - 11：源表不存在 - 12：目标表不存在 - 13：原表和目标表都不存在 - 14：源数据库连接失败 - 15：目标库数据库连接失败 - 16：源数据库执行SQL超时 - 17：目标数据库执行SQL超时 - 18：源数据库执行SQL错误 - 19：目标数据库执行SQL错误 - 20：源库和目标库都不存在 - 21：源库不存在 - 22：目标库不存在 - 23：行数为亿行，未进行对比 - 27：超时
	Status *int32 `json:"status,omitempty"`

	// 行过滤配置条件
	CompareLineConfigFilter *string `json:"compare_line_config_filter,omitempty"`
}

func (o TableLineCompareResultInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableLineCompareResultInfo struct{}"
	}

	return strings.Join([]string{"TableLineCompareResultInfo", string(data)}, " ")
}
