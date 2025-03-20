package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// InterruptRecord 中断记录
type InterruptRecord struct {

	// 记录ID
	Id *string `json:"id,omitempty"`

	// 来源 SLI SLI指标 WARROOM warroom ALERTS 告警 MALFUNCTIONS 故障 OTHERS 其他
	Sources *InterruptRecordSources `json:"sources,omitempty"`

	// 来源ID
	SourcesId *string `json:"sources_id,omitempty"`

	// 来源名称
	SourcesName *string `json:"sources_name,omitempty"`

	// region
	Region *string `json:"region,omitempty"`

	// 不可用开始时间
	UnavailableStartTime *int64 `json:"unavailable_start_time,omitempty"`

	// 不可用结束时间
	UnavailableEndTime *int64 `json:"unavailable_end_time,omitempty"`

	// 描述
	Descriptions *string `json:"descriptions,omitempty"`

	// 创建人Id
	CreatorId *string `json:"creator_id,omitempty"`

	// 创建人
	Creator *string `json:"creator,omitempty"`

	// 修改人ID
	ModifiedId *string `json:"modified_id,omitempty"`

	// 修改人
	ModifiedBy *string `json:"modified_by,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 修改时间
	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o InterruptRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InterruptRecord struct{}"
	}

	return strings.Join([]string{"InterruptRecord", string(data)}, " ")
}

type InterruptRecordSources struct {
	value string
}

type InterruptRecordSourcesEnum struct {
	SLI          InterruptRecordSources
	WARROOM      InterruptRecordSources
	ALERTS       InterruptRecordSources
	MALFUNCTIONS InterruptRecordSources
	OTHERS       InterruptRecordSources
}

func GetInterruptRecordSourcesEnum() InterruptRecordSourcesEnum {
	return InterruptRecordSourcesEnum{
		SLI: InterruptRecordSources{
			value: "SLI",
		},
		WARROOM: InterruptRecordSources{
			value: "WARROOM",
		},
		ALERTS: InterruptRecordSources{
			value: "ALERTS",
		},
		MALFUNCTIONS: InterruptRecordSources{
			value: "MALFUNCTIONS",
		},
		OTHERS: InterruptRecordSources{
			value: "OTHERS",
		},
	}
}

func (c InterruptRecordSources) Value() string {
	return c.value
}

func (c InterruptRecordSources) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InterruptRecordSources) UnmarshalJSON(b []byte) error {
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
