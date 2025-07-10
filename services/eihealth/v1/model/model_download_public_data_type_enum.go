package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DownloadPublicDataTypeEnum **参数解释**：   下载类型，仅支持PUBLIC、EXTRANET。   **约束限制**：   不涉及   **取值范围**：   * PUBLIC：公共数据   * EXTRANET：外部数据   **默认取值**：   不涉及
type DownloadPublicDataTypeEnum struct {
	value string
}

type DownloadPublicDataTypeEnumEnum struct {
	PUBLIC   DownloadPublicDataTypeEnum
	EXTRANET DownloadPublicDataTypeEnum
}

func GetDownloadPublicDataTypeEnumEnum() DownloadPublicDataTypeEnumEnum {
	return DownloadPublicDataTypeEnumEnum{
		PUBLIC: DownloadPublicDataTypeEnum{
			value: "PUBLIC",
		},
		EXTRANET: DownloadPublicDataTypeEnum{
			value: "EXTRANET",
		},
	}
}

func (c DownloadPublicDataTypeEnum) Value() string {
	return c.value
}

func (c DownloadPublicDataTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DownloadPublicDataTypeEnum) UnmarshalJSON(b []byte) error {
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
