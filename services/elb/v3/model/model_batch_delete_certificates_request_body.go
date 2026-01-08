package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteCertificatesRequestBody This is a auto create Body Object
type BatchDeleteCertificatesRequestBody struct {

	// 待删除的证书id列表。
	Certificates []string `json:"certificates"`
}

func (o BatchDeleteCertificatesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteCertificatesRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteCertificatesRequestBody", string(data)}, " ")
}
