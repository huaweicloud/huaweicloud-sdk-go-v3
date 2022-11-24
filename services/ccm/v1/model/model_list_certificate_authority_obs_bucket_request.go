package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListCertificateAuthorityObsBucketRequest struct {
}

func (o ListCertificateAuthorityObsBucketRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCertificateAuthorityObsBucketRequest struct{}"
	}

	return strings.Join([]string{"ListCertificateAuthorityObsBucketRequest", string(data)}, " ")
}
