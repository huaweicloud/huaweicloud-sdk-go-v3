package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// QuotaResourceTypeEnum 配额-资源类型 * GPU_INSTANCES：GPU资源实例数，单位个 * INSTANCES：普通实例数，单位个 * VOLUME_GIGABYTES：磁盘总容量，单位GB * VOLUMES：磁盘数量，单位个 * CORES：CPU数量，单位个 * MEMORY：内存容量，单位MB
type QuotaResourceTypeEnum struct {
	value string
}

type QuotaResourceTypeEnumEnum struct {
	GPU_INSTANCES    QuotaResourceTypeEnum
	INSTANCES        QuotaResourceTypeEnum
	VOLUME_GIGABYTES QuotaResourceTypeEnum
	VOLUMES          QuotaResourceTypeEnum
	CORES            QuotaResourceTypeEnum
	MEMORY           QuotaResourceTypeEnum
}

func GetQuotaResourceTypeEnumEnum() QuotaResourceTypeEnumEnum {
	return QuotaResourceTypeEnumEnum{
		GPU_INSTANCES: QuotaResourceTypeEnum{
			value: "GPU_INSTANCES",
		},
		INSTANCES: QuotaResourceTypeEnum{
			value: "INSTANCES",
		},
		VOLUME_GIGABYTES: QuotaResourceTypeEnum{
			value: "VOLUME_GIGABYTES",
		},
		VOLUMES: QuotaResourceTypeEnum{
			value: "VOLUMES",
		},
		CORES: QuotaResourceTypeEnum{
			value: "CORES",
		},
		MEMORY: QuotaResourceTypeEnum{
			value: "MEMORY",
		},
	}
}

func (c QuotaResourceTypeEnum) Value() string {
	return c.value
}

func (c QuotaResourceTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QuotaResourceTypeEnum) UnmarshalJSON(b []byte) error {
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
