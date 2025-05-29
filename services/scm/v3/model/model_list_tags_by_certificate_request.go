package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTagsByCertificateRequest Request Object
type ListTagsByCertificateRequest struct {

	// 证书id。
	ResourceId string `json:"resource_id"`
}

func (o ListTagsByCertificateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagsByCertificateRequest struct{}"
	}

	return strings.Join([]string{"ListTagsByCertificateRequest", string(data)}, " ")
}
