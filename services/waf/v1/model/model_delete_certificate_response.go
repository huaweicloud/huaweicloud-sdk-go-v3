package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCertificateResponse Response Object
type DeleteCertificateResponse struct {

	// 证书ID
	Id *string `json:"id,omitempty"`

	// 证书名
	Name *string `json:"name,omitempty"`

	// 证书文件，PEM编码
	Content *string `json:"content,omitempty"`

	// 证书私钥，PEM编码
	Key *string `json:"key,omitempty"`

	// 证书过期时间戳
	ExpireTime *int64 `json:"expire_time,omitempty"`

	// 证书过期状态，0-未过期，1-已过期，2-即将过期
	ExpStatus *int32 `json:"exp_status,omitempty"`

	// 证书上传时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 证书关联的域名信息
	BindHost       *[]BindHost `json:"bind_host,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o DeleteCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCertificateResponse struct{}"
	}

	return strings.Join([]string{"DeleteCertificateResponse", string(data)}, " ")
}
