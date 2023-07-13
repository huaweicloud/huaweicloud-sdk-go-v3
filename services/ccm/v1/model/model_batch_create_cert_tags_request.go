package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateCertTagsRequest Request Object
type BatchCreateCertTagsRequest struct {

	// 所需要批量创建标签的证书ID。
	CertificateId string `json:"certificate_id"`

	Body *BatchOperateTagRequestBody `json:"body,omitempty"`
}

func (o BatchCreateCertTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateCertTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateCertTagsRequest", string(data)}, " ")
}
