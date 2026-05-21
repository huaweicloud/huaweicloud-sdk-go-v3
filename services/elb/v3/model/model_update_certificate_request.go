package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCertificateRequest Request Object
type UpdateCertificateRequest struct {

	// **参数解释**：证书ID。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	CertificateId string `json:"certificate_id"`

	Body *UpdateCertificateRequestBody `json:"body,omitempty"`
}

func (o UpdateCertificateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCertificateRequest struct{}"
	}

	return strings.Join([]string{"UpdateCertificateRequest", string(data)}, " ")
}
