package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// InconsistentTypeEnum 通过该类型过滤出与桌面池规格类型不一致的桌面 - product_id: 查找productID与桌面池套餐ID不一致的桌面 - image_id: 查找imageID与桌面池镜像ID不一致的桌面 - disk_num: 查找数据盘数量与桌面池配置不一致的桌面 - disk_size: 查找磁盘累加容量与桌面池配置不一致的桌面
type InconsistentTypeEnum struct {
	value string
}

type InconsistentTypeEnumEnum struct {
	PRODUCT_ID InconsistentTypeEnum
	IMAGE_ID   InconsistentTypeEnum
	DISK_NUM   InconsistentTypeEnum
	DISK_SIZE  InconsistentTypeEnum
}

func GetInconsistentTypeEnumEnum() InconsistentTypeEnumEnum {
	return InconsistentTypeEnumEnum{
		PRODUCT_ID: InconsistentTypeEnum{
			value: "product_id",
		},
		IMAGE_ID: InconsistentTypeEnum{
			value: "image_id",
		},
		DISK_NUM: InconsistentTypeEnum{
			value: "disk_num",
		},
		DISK_SIZE: InconsistentTypeEnum{
			value: "disk_size",
		},
	}
}

func (c InconsistentTypeEnum) Value() string {
	return c.value
}

func (c InconsistentTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InconsistentTypeEnum) UnmarshalJSON(b []byte) error {
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
