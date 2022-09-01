package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CertificatesRspDto struct {

	// 设备CA证书ID，在上传设备CA证书时由平台分配的唯一标识。
	CertificateId *string `json:"certificate_id,omitempty" xml:"certificate_id"`

	// CA证书CN名称。
	CnName *string `json:"cn_name,omitempty" xml:"cn_name"`

	// CA证书所有者。
	Owner *string `json:"owner,omitempty" xml:"owner"`

	// CA证书验证状态。true代表证书已通过验证，可进行设备证书认证接入。false代表证书未通过验证。
	Status *bool `json:"status,omitempty" xml:"status"`

	// CA证书验证码。
	VerifyCode *string `json:"verify_code,omitempty" xml:"verify_code"`

	// 创建证书日期。格式：yyyyMMdd'T'HHmmss'Z'，如20151212T121212Z。
	CreateDate *string `json:"create_date,omitempty" xml:"create_date"`

	// CA证书生效日期。格式：yyyyMMdd'T'HHmmss'Z'，如20151212T121212Z。
	EffectiveDate *string `json:"effective_date,omitempty" xml:"effective_date"`

	// CA证书失效日期。格式：yyyyMMdd'T'HHmmss'Z'，如20151212T121212Z。
	ExpiryDate *string `json:"expiry_date,omitempty" xml:"expiry_date"`
}

func (o CertificatesRspDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CertificatesRspDto struct{}"
	}

	return strings.Join([]string{"CertificatesRspDto", string(data)}, " ")
}
