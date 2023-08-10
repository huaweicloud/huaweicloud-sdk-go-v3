package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DrugFileSource 受体的数据源：外部网络数据（如RCSB在线数据库）、用户私有数据中心、承载租户公共数据（含样例/公共库）
type DrugFileSource struct {
	value string
}

type DrugFileSourceEnum struct {
	EXTRANET DrugFileSource
	PRIVATE  DrugFileSource
	PUBLIC   DrugFileSource
	RAW      DrugFileSource
}

func GetDrugFileSourceEnum() DrugFileSourceEnum {
	return DrugFileSourceEnum{
		EXTRANET: DrugFileSource{
			value: "EXTRANET",
		},
		PRIVATE: DrugFileSource{
			value: "PRIVATE",
		},
		PUBLIC: DrugFileSource{
			value: "PUBLIC",
		},
		RAW: DrugFileSource{
			value: "RAW",
		},
	}
}

func (c DrugFileSource) Value() string {
	return c.value
}

func (c DrugFileSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DrugFileSource) UnmarshalJSON(b []byte) error {
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
