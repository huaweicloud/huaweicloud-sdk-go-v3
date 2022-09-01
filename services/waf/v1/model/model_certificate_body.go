package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CertificateBody struct {

	// 证书id
	Id string `json:"id" xml:"id"`

	// 证书名
	Name string `json:"name" xml:"name"`

	// 证书过期时间戳
	ExpireTime *int64 `json:"expire_time,omitempty" xml:"expire_time"`

	// 证书过期状态，0-未过期，1-已过期，2-即将过期（一个月内即将过期）
	ExpStatus *int32 `json:"exp_status,omitempty" xml:"exp_status"`

	// 证书上传时间戳
	Timestamp int64 `json:"timestamp" xml:"timestamp"`

	// 证书关联的域名信息
	BindHost *[]BindHost `json:"bind_host,omitempty" xml:"bind_host"`
}

func (o CertificateBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CertificateBody struct{}"
	}

	return strings.Join([]string{"CertificateBody", string(data)}, " ")
}
