package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateCertificateResponse struct {

	// 证书ID
	Id *string `json:"id,omitempty" xml:"id"`

	// 证书名
	Name *string `json:"name,omitempty" xml:"name"`

	// 证书文件，PEM编码
	Content *string `json:"content,omitempty" xml:"content"`

	// 证书私钥，PEM编码
	Key *string `json:"key,omitempty" xml:"key"`

	// 证书过期时间戳
	ExpireTime *int64 `json:"expire_time,omitempty" xml:"expire_time"`

	// 证书过期状态，0-未过期，1-已过期，2-即将过期
	ExpStatus *int32 `json:"exp_status,omitempty" xml:"exp_status"`

	// 证书上传时间戳
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp"`

	// 证书关联的域名信息
	BindHost       *[]BindHost `json:"bind_host,omitempty" xml:"bind_host"`
	HttpStatusCode int         `json:"-"`
}

func (o CreateCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertificateResponse struct{}"
	}

	return strings.Join([]string{"CreateCertificateResponse", string(data)}, " ")
}
