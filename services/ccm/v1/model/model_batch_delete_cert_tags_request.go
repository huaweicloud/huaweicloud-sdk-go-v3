package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteCertTagsRequest Request Object
type BatchDeleteCertTagsRequest struct {

	// 所需要批量删除标签的证书ID。
	CertificateId string `json:"certificate_id"`

	Body *BatchOperateTagRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteCertTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteCertTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteCertTagsRequest", string(data)}, " ")
}
