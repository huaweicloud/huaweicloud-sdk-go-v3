package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type DeleteDatabaseRequest struct {
	// DDM实例ID。

	InstanceId string `json:"instance_id"`
	// 需要查询的逻辑库名称，不区分大小写。

	DdmDbname string `json:"ddm_dbname"`
	// 是否同时删除关联后端数据库实例上存储的数据。 - 取值为“true”：删除。 - 取值为空或“false”：不删除。 默认值为空。

	DeleteRdsData *DeleteDatabaseRequestDeleteRdsData `json:"delete_rds_data,omitempty"`
}

func (o DeleteDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDatabaseRequest struct{}"
	}

	return strings.Join([]string{"DeleteDatabaseRequest", string(data)}, " ")
}

type DeleteDatabaseRequestDeleteRdsData struct {
	value string
}

type DeleteDatabaseRequestDeleteRdsDataEnum struct {
	TRUE  DeleteDatabaseRequestDeleteRdsData
	FALSE DeleteDatabaseRequestDeleteRdsData
}

func GetDeleteDatabaseRequestDeleteRdsDataEnum() DeleteDatabaseRequestDeleteRdsDataEnum {
	return DeleteDatabaseRequestDeleteRdsDataEnum{
		TRUE: DeleteDatabaseRequestDeleteRdsData{
			value: "true",
		},
		FALSE: DeleteDatabaseRequestDeleteRdsData{
			value: "false",
		},
	}
}

func (c DeleteDatabaseRequestDeleteRdsData) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteDatabaseRequestDeleteRdsData) UnmarshalJSON(b []byte) error {
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
