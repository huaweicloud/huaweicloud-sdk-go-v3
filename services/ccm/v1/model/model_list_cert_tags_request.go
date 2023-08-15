package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCertTagsRequest Request Object
type ListCertTagsRequest struct {

	// 所需要查询标签列表的证书ID。
	CertificateId string `json:"certificate_id"`
}

func (o ListCertTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCertTagsRequest struct{}"
	}

	return strings.Join([]string{"ListCertTagsRequest", string(data)}, " ")
}
