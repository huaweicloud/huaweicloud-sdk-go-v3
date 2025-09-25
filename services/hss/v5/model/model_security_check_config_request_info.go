package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SecurityCheckConfigRequestInfo 待修改的安全体检定时配置信息
type SecurityCheckConfigRequestInfo struct {

	// 定时任务ID，前台只保存，不显示
	TaskId *int64 `json:"task_id,omitempty"`

	// 体检状态，是否开启
	Status *bool `json:"status,omitempty"`

	// 体检周期类型：day-按天，week-按周
	CheckPeriodType *SecurityCheckConfigRequestInfoCheckPeriodType `json:"check_period_type,omitempty"`

	// 按天的周期
	DayPeriod *int32 `json:"day_period,omitempty"`

	// 按周的周期,选中的日期放入此列表，取值包含：mon,tue,wed,thu,fri,sat,sun
	WeekPeriod *[]string `json:"week_period,omitempty"`

	// 体检时间：小时
	Hour *int32 `json:"hour,omitempty"`

	// 体检内容，取值包含：asset,vul,baseline,event
	Content *[]string `json:"content,omitempty"`

	// 已选服务器ID列表
	HostIdList *[]string `json:"host_id_list,omitempty"`

	// 服务器是否选择全部，全选跟查询条件无关
	OperateAll *bool `json:"operate_all,omitempty"`
}

func (o SecurityCheckConfigRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityCheckConfigRequestInfo struct{}"
	}

	return strings.Join([]string{"SecurityCheckConfigRequestInfo", string(data)}, " ")
}

type SecurityCheckConfigRequestInfoCheckPeriodType struct {
	value string
}

type SecurityCheckConfigRequestInfoCheckPeriodTypeEnum struct {
	DAY  SecurityCheckConfigRequestInfoCheckPeriodType
	WEEK SecurityCheckConfigRequestInfoCheckPeriodType
}

func GetSecurityCheckConfigRequestInfoCheckPeriodTypeEnum() SecurityCheckConfigRequestInfoCheckPeriodTypeEnum {
	return SecurityCheckConfigRequestInfoCheckPeriodTypeEnum{
		DAY: SecurityCheckConfigRequestInfoCheckPeriodType{
			value: "day",
		},
		WEEK: SecurityCheckConfigRequestInfoCheckPeriodType{
			value: "week",
		},
	}
}

func (c SecurityCheckConfigRequestInfoCheckPeriodType) Value() string {
	return c.value
}

func (c SecurityCheckConfigRequestInfoCheckPeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SecurityCheckConfigRequestInfoCheckPeriodType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
