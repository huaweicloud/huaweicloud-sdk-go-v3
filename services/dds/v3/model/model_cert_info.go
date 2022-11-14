package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CertInfo struct {

	// 证书种类
	Category CertInfoCategory `json:"category"`

	// 证书下载链接
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
	NATIONAL      CertInfoCategory
}

func GetCertInfoCategoryEnum() CertInfoCategoryEnum {
	return CertInfoCategoryEnum{
		INTERNATIONAL: CertInfoCategory{
			value: "international",
		},
		NATIONAL: CertInfoCategory{
			value: "national",
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
