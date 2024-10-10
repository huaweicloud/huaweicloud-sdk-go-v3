package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CertInfo struct {

	// 证书名称
	CertName string `json:"cert_name"`

	// 证书id
	Id string `json:"id"`

	// 适用域名
	ApplyDomain string `json:"apply_domain"`

	// 过期时间
	ExpireTime int64 `json:"expire_time"`

	// 过期状态
	ExpireStatus int32 `json:"expire_status"`
}

func (o CertInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CertInfo struct{}"
	}

	return strings.Join([]string{"CertInfo", string(data)}, " ")
}
