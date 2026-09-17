package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SlowLogTplContrast 慢日志模板对比结果
type SlowLogTplContrast struct {

	// 前一日慢日志模板数据列表
	TemplateOfPreDay *[]SlowSqlTemplate `json:"template_of_pre_day,omitempty"`

	// 当日慢日志模板数据列表
	TemplateOfCurDay *[]SlowSqlTemplate `json:"template_of_cur_day,omitempty"`

	// 执行耗时是否增长
	ExecuteTimeIncrease *bool `json:"execute_time_increase,omitempty"`

	// 锁等待耗时是否增长
	LockWaitIncrease *bool `json:"lock_wait_increase,omitempty"`

	// 是否新增模板
	NewTemplate *bool `json:"new_template,omitempty"`
}

func (o SlowLogTplContrast) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowLogTplContrast struct{}"
	}

	return strings.Join([]string{"SlowLogTplContrast", string(data)}, " ")
}
