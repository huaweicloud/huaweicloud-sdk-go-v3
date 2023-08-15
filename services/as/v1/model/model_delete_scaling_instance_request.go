package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteScalingInstanceRequest Request Object
type DeleteScalingInstanceRequest struct {

	// 实例ID，可参考[查询弹性伸缩组中的实例列表](https://support.huaweicloud.com/api-as/as_06_0301.html)获取。
	InstanceId string `json:"instance_id"`

	// 实例移出伸缩组，是否删除云服务器实例。默认为no；可选值为yes或no。
	InstanceDelete *DeleteScalingInstanceRequestInstanceDelete `json:"instance_delete,omitempty"`
}

func (o DeleteScalingInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScalingInstanceRequest struct{}"
	}

	return strings.Join([]string{"DeleteScalingInstanceRequest", string(data)}, " ")
}

type DeleteScalingInstanceRequestInstanceDelete struct {
	value string
}

type DeleteScalingInstanceRequestInstanceDeleteEnum struct {
	YES DeleteScalingInstanceRequestInstanceDelete
	NO  DeleteScalingInstanceRequestInstanceDelete
}

func GetDeleteScalingInstanceRequestInstanceDeleteEnum() DeleteScalingInstanceRequestInstanceDeleteEnum {
	return DeleteScalingInstanceRequestInstanceDeleteEnum{
		YES: DeleteScalingInstanceRequestInstanceDelete{
			value: "yes",
		},
		NO: DeleteScalingInstanceRequestInstanceDelete{
			value: "no",
		},
	}
}

func (c DeleteScalingInstanceRequestInstanceDelete) Value() string {
	return c.value
}

func (c DeleteScalingInstanceRequestInstanceDelete) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteScalingInstanceRequestInstanceDelete) UnmarshalJSON(b []byte) error {
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
