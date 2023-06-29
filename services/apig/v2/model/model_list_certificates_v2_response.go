package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCertificatesV2Response Response Object
type ListCertificatesV2Response struct {

	// 本次返回的列表长度
	Size int32 `json:"size"`

	// 满足条件的记录数
	Total int64 `json:"total"`

	// 证书基本内容
	Certs          *[]CertBase `json:"certs,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListCertificatesV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCertificatesV2Response struct{}"
	}

	return strings.Join([]string{"ListCertificatesV2Response", string(data)}, " ")
}
