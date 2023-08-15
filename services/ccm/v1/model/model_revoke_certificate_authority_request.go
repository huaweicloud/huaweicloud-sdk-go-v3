package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RevokeCertificateAuthorityRequest Request Object
type RevokeCertificateAuthorityRequest struct {

	// 所要吊销的子CA ID。
	CaId string `json:"ca_id"`

	Body *RevokeCertificateRequestBody `json:"body,omitempty"`
}

func (o RevokeCertificateAuthorityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RevokeCertificateAuthorityRequest struct{}"
	}

	return strings.Join([]string{"RevokeCertificateAuthorityRequest", string(data)}, " ")
}
