package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Resources struct {

	// 函数配额限制。
	Quota *int32 `json:"quota,omitempty"`

	// 已使用的配额。
	Used *int32 `json:"used,omitempty"`

	// 资源类型。 fgs_func_scale_down_timeout：v1版本函数的实例闲置释放时间 fgs_func_occurs：v1版本函数为实例数配额, v2版本函数为预留实例配额 fgs_func_pat_idle_time：vpc函数的pat闲置释放时间 fgs_func_num：用户函数数量配额 fgs_func_code_size：用户函数总代码大小配额 fgs_workflow_num：用户函数流数量配额 fgs_on_demand_instance_limit：v2版本函数单函数最大实例数配额 fgs_func_qos_limit 用户函数实例数配额
	Type *ResourcesType `json:"type,omitempty"`

	// 资源的计数单位。fgs_func_code_size,单位为MB,其他场景无单位
	Unit *string `json:"unit,omitempty"`
}

func (o Resources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resources struct{}"
	}

	return strings.Join([]string{"Resources", string(data)}, " ")
}

type ResourcesType struct {
	value string
}

type ResourcesTypeEnum struct {
	FGS_FUNC_SCALE_DOWN_TIMEOUT  ResourcesType
	FGS_FUNC_OCCURS              ResourcesType
	FGS_FUNC_PAT_IDLE_TIME       ResourcesType
	FGS_FUNC_NUM                 ResourcesType
	FGS_FUNC_CODE_SIZE           ResourcesType
	FGS_WORKFLOW_NUM             ResourcesType
	FGS_ON_DEMAND_INSTANCE_LIMIT ResourcesType
	FGS_FUNC_QOS_LIMIT           ResourcesType
}

func GetResourcesTypeEnum() ResourcesTypeEnum {
	return ResourcesTypeEnum{
		FGS_FUNC_SCALE_DOWN_TIMEOUT: ResourcesType{
			value: "fgs_func_scale_down_timeout",
		},
		FGS_FUNC_OCCURS: ResourcesType{
			value: "fgs_func_occurs",
		},
		FGS_FUNC_PAT_IDLE_TIME: ResourcesType{
			value: "fgs_func_pat_idle_time",
		},
		FGS_FUNC_NUM: ResourcesType{
			value: "fgs_func_num",
		},
		FGS_FUNC_CODE_SIZE: ResourcesType{
			value: "fgs_func_code_size",
		},
		FGS_WORKFLOW_NUM: ResourcesType{
			value: "fgs_workflow_num",
		},
		FGS_ON_DEMAND_INSTANCE_LIMIT: ResourcesType{
			value: "fgs_on_demand_instance_limit",
		},
		FGS_FUNC_QOS_LIMIT: ResourcesType{
			value: "fgs_func_qos_limit",
		},
	}
}

func (c ResourcesType) Value() string {
	return c.value
}

func (c ResourcesType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourcesType) UnmarshalJSON(b []byte) error {
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
