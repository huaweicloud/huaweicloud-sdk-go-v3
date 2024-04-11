package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataCompareOverviewInfo 行数对比总览详情
type DataCompareOverviewInfo struct {

	// 源库库名
	SourceDbName *string `json:"source_db_name,omitempty"`

	// 目标库库名
	TargetDbName *string `json:"target_db_name,omitempty"`

	// 对比状态。 - 0：对比不一致 - 2：对比一致 - 3：目标库表不存在 - 4：对比失败 - 5：正在对比中 - 6：等待对比中 - 7：任务已取消 - 8：源库为空 - 9：目标库为空 - 10：源库和目标库都为空 - 11：源表不存在 - 12：目标表不存在 - 13：原表和目标表都不存在 - 14：源数据库连接失败 - 15：目标库数据库连接失败 - 16：源数据库执行SQL超时 - 17：目标数据库执行SQL超时 - 18：源数据库执行SQL错误 - 19：目标数据库执行SQL错误 - 20：源库和目标库都不存在 - 21：源库不存在 - 22：目标库不存在 - 23：行数为亿行，未进行对比 - 27：超时
	Status *int32 `json:"status,omitempty"`
}

func (o DataCompareOverviewInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataCompareOverviewInfo struct{}"
	}

	return strings.Join([]string{"DataCompareOverviewInfo", string(data)}, " ")
}
