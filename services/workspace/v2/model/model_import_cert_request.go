package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportCertRequest Request Object
type ImportCertRequest struct {
	Body *ImportCertificatePemReq `json:"body,omitempty"`
}

func (o ImportCertRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportCertRequest struct{}"
	}

	return strings.Join([]string{"ImportCertRequest", string(data)}, " ")
}
