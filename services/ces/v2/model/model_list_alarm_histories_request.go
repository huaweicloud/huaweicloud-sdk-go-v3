package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAlarmHistoriesRequest Request Object
type ListAlarmHistoriesRequest struct {

	// 告警ID,以al开头，后跟22位由字母或数字组成的字符串
	AlarmId *string `json:"alarm_id,omitempty"`

	// 告警规则名称
	Name *string `json:"name,omitempty"`

	// 告警规则状态, ok为正常，alarm为告警，invalid为已失效
	Status *string `json:"status,omitempty"`

	// 告警级别, 1为紧急，2为重要，3为次要，4为提示
	Level *int32 `json:"level,omitempty"`

	// 查询服务的命名空间，各服务命名空间请参考[服务命名空间](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)
	Namespace *string `json:"namespace,omitempty"`

	// 告警资源ID，多维度情况按字母升序排列并使用逗号分隔
	ResourceId *string `json:"resource_id,omitempty"`

	// 查询告警记录的起始时间，例如：2022-02-10T10:05:46+08:00
	From *string `json:"from,omitempty"`

	// 查询告警记录的截止时间，例如：2022-02-10T10:05:47+08:00
	To *string `json:"to,omitempty"`

	// 分页偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 分页大小
	Limit *int32 `json:"limit,omitempty"`

	// 按关键字排序, 默认为update_time, {first_alarm_time: 告警产生时间, update_time: 更新时间, alarm_level: 告警级别, record_id：表记录主键}
	OrderBy *ListAlarmHistoriesRequestOrderBy `json:"order_by,omitempty"`
}

func (o ListAlarmHistoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmHistoriesRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmHistoriesRequest", string(data)}, " ")
}

type ListAlarmHistoriesRequestOrderBy struct {
	value string
}

type ListAlarmHistoriesRequestOrderByEnum struct {
	FIRST_ALARM_TIME ListAlarmHistoriesRequestOrderBy
	UPDATE_TIME      ListAlarmHistoriesRequestOrderBy
	ALARM_LEVEL      ListAlarmHistoriesRequestOrderBy
	RECORD_ID        ListAlarmHistoriesRequestOrderBy
}

func GetListAlarmHistoriesRequestOrderByEnum() ListAlarmHistoriesRequestOrderByEnum {
	return ListAlarmHistoriesRequestOrderByEnum{
		FIRST_ALARM_TIME: ListAlarmHistoriesRequestOrderBy{
			value: "first_alarm_time",
		},
		UPDATE_TIME: ListAlarmHistoriesRequestOrderBy{
			value: "update_time",
		},
		ALARM_LEVEL: ListAlarmHistoriesRequestOrderBy{
			value: "alarm_level",
		},
		RECORD_ID: ListAlarmHistoriesRequestOrderBy{
			value: "record_id",
		},
	}
}

func (c ListAlarmHistoriesRequestOrderBy) Value() string {
	return c.value
}

func (c ListAlarmHistoriesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAlarmHistoriesRequestOrderBy) UnmarshalJSON(b []byte) error {
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
