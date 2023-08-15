package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCertificateRequestBody This is a auto create Body Object
type UpdateCertificateRequestBody struct {
	Certificate *UpdateCertificateOption `json:"certificate"`
}

func (o UpdateCertificateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCertificateRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateCertificateRequestBody", string(data)}, " ")
}
