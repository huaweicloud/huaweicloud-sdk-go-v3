package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SqlTplCmp SQL模板对比
type SqlTplCmp struct {
	SqlTplDto1 *TplCmp `json:"sql_tpl_dto1,omitempty"`

	SqlTplDto2 *TplCmp `json:"sql_tpl_dto2,omitempty"`

	// 是否为新增数据
	New *bool `json:"new,omitempty"`

	// 是否有执行时间增长
	ExecuteTimeIncrease *bool `json:"execute_time_increase,omitempty"`

	// 是否有锁等待时间增长
	LockWaitIncrease *bool `json:"lock_wait_increase,omitempty"`
}

func (o SqlTplCmp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlTplCmp struct{}"
	}

	return strings.Join([]string{"SqlTplCmp", string(data)}, " ")
}
