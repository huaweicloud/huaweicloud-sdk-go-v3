package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteInstanceRequest Request Object
type DeleteInstanceRequest struct {

	// DDM实例ID。
	InstanceId string `json:"instance_id"`

	// 是否同时删除关联后端数据库实例上存储的数据。  - 取值为空或“true”：删除。 - 取值为“false”：不删除。 默认值为空。
	DeleteRdsData *DeleteInstanceRequestDeleteRdsData `json:"delete_rds_data,omitempty"`
}

func (o DeleteInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceRequest struct{}"
	}

	return strings.Join([]string{"DeleteInstanceRequest", string(data)}, " ")
}

type DeleteInstanceRequestDeleteRdsData struct {
	value string
}

type DeleteInstanceRequestDeleteRdsDataEnum struct {
	TRUE  DeleteInstanceRequestDeleteRdsData
	FALSE DeleteInstanceRequestDeleteRdsData
}

func GetDeleteInstanceRequestDeleteRdsDataEnum() DeleteInstanceRequestDeleteRdsDataEnum {
	return DeleteInstanceRequestDeleteRdsDataEnum{
		TRUE: DeleteInstanceRequestDeleteRdsData{
			value: "true",
		},
		FALSE: DeleteInstanceRequestDeleteRdsData{
			value: "false",
		},
	}
}

func (c DeleteInstanceRequestDeleteRdsData) Value() string {
	return c.value
}

func (c DeleteInstanceRequestDeleteRdsData) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteInstanceRequestDeleteRdsData) UnmarshalJSON(b []byte) error {
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
