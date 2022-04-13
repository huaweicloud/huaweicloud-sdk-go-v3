package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type DownloadBlockchainCertRequest struct {
	// blockchainID

	BlockchainId string `json:"blockchain_id"`
	// order或者peer组织名称

	OrgName string `json:"org_name"`
	// 下载证书类别

	CertType DownloadBlockchainCertRequestCertType `json:"cert_type"`
}

func (o DownloadBlockchainCertRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadBlockchainCertRequest struct{}"
	}

	return strings.Join([]string{"DownloadBlockchainCertRequest", string(data)}, " ")
}

type DownloadBlockchainCertRequestCertType struct {
	value string
}

type DownloadBlockchainCertRequestCertTypeEnum struct {
	ADMIN DownloadBlockchainCertRequestCertType
	USER  DownloadBlockchainCertRequestCertType
	CA    DownloadBlockchainCertRequestCertType
}

func GetDownloadBlockchainCertRequestCertTypeEnum() DownloadBlockchainCertRequestCertTypeEnum {
	return DownloadBlockchainCertRequestCertTypeEnum{
		ADMIN: DownloadBlockchainCertRequestCertType{
			value: "admin",
		},
		USER: DownloadBlockchainCertRequestCertType{
			value: "user",
		},
		CA: DownloadBlockchainCertRequestCertType{
			value: "ca",
		},
	}
}

func (c DownloadBlockchainCertRequestCertType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DownloadBlockchainCertRequestCertType) UnmarshalJSON(b []byte) error {
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
