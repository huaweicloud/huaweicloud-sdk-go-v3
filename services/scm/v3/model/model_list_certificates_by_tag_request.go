package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCertificatesByTagRequest Request Object
type ListCertificatesByTagRequest struct {

	// 定值为resource_instances。
	ResourceInstances string `json:"resource_instances"`

	Body *ListCertificatesByTagRequestBody `json:"body,omitempty"`
}

func (o ListCertificatesByTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCertificatesByTagRequest struct{}"
	}

	return strings.Join([]string{"ListCertificatesByTagRequest", string(data)}, " ")
}
