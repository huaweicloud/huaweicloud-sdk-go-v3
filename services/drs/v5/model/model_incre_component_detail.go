package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// IncreComponentDetail 增量组件详情。
type IncreComponentDetail struct {

	// 组件类型 - capture：抓取 - apply：回放
	Type *IncreComponentDetailType `json:"type,omitempty"`

	// 状态。 - STOPPED：停止 - STARTED：运行中 - STOPPING：停止中 - STARTING：启动中
	Status *string `json:"status,omitempty"`

	// 启动时间
	StartTime *string `json:"start_time,omitempty"`

	// 启动位点
	StartPoint *string `json:"start_point,omitempty"`

	// 当前位点
	CurrentPoint *string `json:"current_point,omitempty"`

	// 解析时间
	ResolutionTime *string `json:"resolution_time,omitempty"`

	// 时延，单位：秒
	Delay *string `json:"delay,omitempty"`
}

func (o IncreComponentDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IncreComponentDetail struct{}"
	}

	return strings.Join([]string{"IncreComponentDetail", string(data)}, " ")
}

type IncreComponentDetailType struct {
	value string
}

type IncreComponentDetailTypeEnum struct {
	CAPTURE IncreComponentDetailType
	APPLY   IncreComponentDetailType
}

func GetIncreComponentDetailTypeEnum() IncreComponentDetailTypeEnum {
	return IncreComponentDetailTypeEnum{
		CAPTURE: IncreComponentDetailType{
			value: "capture",
		},
		APPLY: IncreComponentDetailType{
			value: "apply",
		},
	}
}

func (c IncreComponentDetailType) Value() string {
	return c.value
}

func (c IncreComponentDetailType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IncreComponentDetailType) UnmarshalJSON(b []byte) error {
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
