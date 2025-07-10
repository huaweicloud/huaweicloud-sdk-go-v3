package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AutoscalePolicy 弹性伸缩策略。
type AutoscalePolicy struct {

	// 弹性伸缩类型，ACCESS_CREATED：接入时创建，AUTO_CREATED：弹性伸缩。
	AutoscaleType *AutoscalePolicyAutoscaleType `json:"autoscale_type,omitempty"`

	// 自动创建桌面上限。
	MaxAutoCreated *int32 `json:"max_auto_created,omitempty"`

	// 空闲桌面低于多少时开始自动创建桌面。
	MinIdle *int32 `json:"min_idle,omitempty"`

	// 一次自动创建桌面的数量。
	OnceAutoCreated *int32 `json:"once_auto_created,omitempty"`
}

func (o AutoscalePolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoscalePolicy struct{}"
	}

	return strings.Join([]string{"AutoscalePolicy", string(data)}, " ")
}

type AutoscalePolicyAutoscaleType struct {
	value string
}

type AutoscalePolicyAutoscaleTypeEnum struct {
	ACCESS_CREATED AutoscalePolicyAutoscaleType
	AUTO_CREATED   AutoscalePolicyAutoscaleType
}

func GetAutoscalePolicyAutoscaleTypeEnum() AutoscalePolicyAutoscaleTypeEnum {
	return AutoscalePolicyAutoscaleTypeEnum{
		ACCESS_CREATED: AutoscalePolicyAutoscaleType{
			value: "ACCESS_CREATED",
		},
		AUTO_CREATED: AutoscalePolicyAutoscaleType{
			value: "AUTO_CREATED",
		},
	}
}

func (c AutoscalePolicyAutoscaleType) Value() string {
	return c.value
}

func (c AutoscalePolicyAutoscaleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AutoscalePolicyAutoscaleType) UnmarshalJSON(b []byte) error {
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
