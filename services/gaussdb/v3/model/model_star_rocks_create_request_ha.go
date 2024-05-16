package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StarRocksCreateRequestHa 部署信息。
type StarRocksCreateRequestHa struct {

	// 部署模式。
	Mode StarRocksCreateRequestHaMode `json:"mode"`
}

func (o StarRocksCreateRequestHa) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StarRocksCreateRequestHa struct{}"
	}

	return strings.Join([]string{"StarRocksCreateRequestHa", string(data)}, " ")
}

type StarRocksCreateRequestHaMode struct {
	value string
}

type StarRocksCreateRequestHaModeEnum struct {
	SINGLE  StarRocksCreateRequestHaMode
	CLUSTER StarRocksCreateRequestHaMode
}

func GetStarRocksCreateRequestHaModeEnum() StarRocksCreateRequestHaModeEnum {
	return StarRocksCreateRequestHaModeEnum{
		SINGLE: StarRocksCreateRequestHaMode{
			value: "Single",
		},
		CLUSTER: StarRocksCreateRequestHaMode{
			value: "Cluster",
		},
	}
}

func (c StarRocksCreateRequestHaMode) Value() string {
	return c.value
}

func (c StarRocksCreateRequestHaMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StarRocksCreateRequestHaMode) UnmarshalJSON(b []byte) error {
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
