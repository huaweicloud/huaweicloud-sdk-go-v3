package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListCertificateResponse struct {

	// 私有证书总数。
	Total *int32 `json:"total,omitempty"`

	// 证书列表，详情请参见**Certificates**字段数据结构说明。
	Certificates   *[]Certificates `json:"certificates,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCertificateResponse struct{}"
	}

	return strings.Join([]string{"ListCertificateResponse", string(data)}, " ")
}
