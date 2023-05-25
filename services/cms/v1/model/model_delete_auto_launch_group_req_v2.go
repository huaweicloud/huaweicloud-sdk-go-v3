package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// This is a delete Request Object
type DeleteAutoLaunchGroupReqV2 struct {

	// 删除智能购买组时组内实例中断行为。枚举值 terminate：释放，由delete_publicip和delete_volume决定是否释放ip和卷 noTermination：不释放，ip和卷也不释放 默认值：terminate
	DeleteInstances *DeleteAutoLaunchGroupReqV2DeleteInstances `json:"delete_instances,omitempty"`
}

func (o DeleteAutoLaunchGroupReqV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAutoLaunchGroupReqV2 struct{}"
	}

	return strings.Join([]string{"DeleteAutoLaunchGroupReqV2", string(data)}, " ")
}

type DeleteAutoLaunchGroupReqV2DeleteInstances struct {
	value string
}

type DeleteAutoLaunchGroupReqV2DeleteInstancesEnum struct {
	TERMINATE      DeleteAutoLaunchGroupReqV2DeleteInstances
	NO_TERMINATION DeleteAutoLaunchGroupReqV2DeleteInstances
}

func GetDeleteAutoLaunchGroupReqV2DeleteInstancesEnum() DeleteAutoLaunchGroupReqV2DeleteInstancesEnum {
	return DeleteAutoLaunchGroupReqV2DeleteInstancesEnum{
		TERMINATE: DeleteAutoLaunchGroupReqV2DeleteInstances{
			value: "terminate",
		},
		NO_TERMINATION: DeleteAutoLaunchGroupReqV2DeleteInstances{
			value: "noTermination",
		},
	}
}

func (c DeleteAutoLaunchGroupReqV2DeleteInstances) Value() string {
	return c.value
}

func (c DeleteAutoLaunchGroupReqV2DeleteInstances) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteAutoLaunchGroupReqV2DeleteInstances) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
