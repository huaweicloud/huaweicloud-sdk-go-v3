package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCertKeyResponse Response Object
type ShowCertKeyResponse struct {

	// 源端证书
	Cert *string `json:"cert,omitempty"`

	// 源端私钥
	PrivateKey *string `json:"private_key,omitempty"`

	// ca证书
	Ca *string `json:"ca,omitempty"`

	// 目的端管理层证书
	TargetMgmtCert *string `json:"target_mgmt_cert,omitempty"`

	// 目的端管理层私钥
	TargetMgmtPrivateKey *string `json:"target_mgmt_private_key,omitempty"`

	// 目的端数据层证书
	TargetDataCert *string `json:"target_data_cert,omitempty"`

	// 目的端数据层私钥
	TargetDataPrivateKey *string `json:"target_data_private_key,omitempty"`
	HttpStatusCode       int     `json:"-"`
}

func (o ShowCertKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCertKeyResponse struct{}"
	}

	return strings.Join([]string{"ShowCertKeyResponse", string(data)}, " ")
}
