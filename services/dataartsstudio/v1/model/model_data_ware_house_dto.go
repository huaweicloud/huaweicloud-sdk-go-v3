package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DataWareHouseDto struct {

	// 数据连接ID
	DwId *string `json:"dw_id,omitempty"`

	// 数据连接名称
	DwName *string `json:"dw_name,omitempty"`

	// 数据连接类型：   * HIVE数据源   * DWS数据源   * SPARK数据源
	DwType *DataWareHouseDtoDwType `json:"dw_type,omitempty"`

	DwConfig *DataWareHouseDtoDwConfig `json:"dw_config,omitempty"`
}

func (o DataWareHouseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataWareHouseDto struct{}"
	}

	return strings.Join([]string{"DataWareHouseDto", string(data)}, " ")
}

type DataWareHouseDtoDwType struct {
	value string
}

type DataWareHouseDtoDwTypeEnum struct {
	HIVE  DataWareHouseDtoDwType
	DWS   DataWareHouseDtoDwType
	SPARK DataWareHouseDtoDwType
}

func GetDataWareHouseDtoDwTypeEnum() DataWareHouseDtoDwTypeEnum {
	return DataWareHouseDtoDwTypeEnum{
		HIVE: DataWareHouseDtoDwType{
			value: "HIVE",
		},
		DWS: DataWareHouseDtoDwType{
			value: "DWS",
		},
		SPARK: DataWareHouseDtoDwType{
			value: "SPARK",
		},
	}
}

func (c DataWareHouseDtoDwType) Value() string {
	return c.value
}

func (c DataWareHouseDtoDwType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataWareHouseDtoDwType) UnmarshalJSON(b []byte) error {
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
