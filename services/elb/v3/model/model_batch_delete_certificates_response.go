package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteCertificatesResponse Response Object
type BatchDeleteCertificatesResponse struct {

	// 证书批量删除后的响应结果。
	Certificates   *[]BatchDeleteCertificatesResp `json:"certificates,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o BatchDeleteCertificatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteCertificatesResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteCertificatesResponse", string(data)}, " ")
}
