package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteCertificatesRequest Request Object
type BatchDeleteCertificatesRequest struct {
	Body *BatchDeleteCertificatesRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteCertificatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteCertificatesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteCertificatesRequest", string(data)}, " ")
}
