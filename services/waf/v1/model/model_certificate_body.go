package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CertificateBody struct {
	// 证书id

	Id string `json:"id"`
	// 证书名

	Name string `json:"name"`
	// 证书文件，PEM编码

	Content *string `json:"content,omitempty"`
	// 证书私钥，PEM编码

	Key *string `json:"key,omitempty"`
	// 证书过期时间戳

	ExpireTime *int64 `json:"expire_time,omitempty"`
	// 证书过期状态，0-未过期，1-已过期，2-即将过期（一个月内即将过期）

	ExpStatus *int32 `json:"exp_status,omitempty"`
	// 证书上传时间戳

	Timestamp int64 `json:"timestamp"`
	// 证书关联的域名信息

	BindHost *[]BindHost `json:"bind_host,omitempty"`
}

func (o CertificateBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CertificateBody struct{}"
	}

	return strings.Join([]string{"CertificateBody", string(data)}, " ")
}
