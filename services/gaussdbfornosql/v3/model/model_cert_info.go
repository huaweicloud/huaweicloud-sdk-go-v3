package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CertInfo struct {

	// **参数解释：** 证书种类。 **取值范围：** - international：国际证书。
	Category CertInfoCategory `json:"category"`

	// **参数解释：** 证书下载链接。 **取值范围：** 不涉及。
	DownloadLink string `json:"download_link"`
}

func (o CertInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CertInfo struct{}"
	}

	return strings.Join([]string{"CertInfo", string(data)}, " ")
}

type CertInfoCategory struct {
	value string
}

type CertInfoCategoryEnum struct {
	INTERNATIONAL CertInfoCategory
}

func GetCertInfoCategoryEnum() CertInfoCategoryEnum {
	return CertInfoCategoryEnum{
		INTERNATIONAL: CertInfoCategory{
			value: "international",
		},
	}
}

func (c CertInfoCategory) Value() string {
	return c.value
}

func (c CertInfoCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CertInfoCategory) UnmarshalJSON(b []byte) error {
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
