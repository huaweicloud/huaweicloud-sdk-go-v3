package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSpCertificatesResponse Response Object
type ListSpCertificatesResponse struct {
	Body           *[]SpCertificateDto `json:"body,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListSpCertificatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSpCertificatesResponse struct{}"
	}

	return strings.Join([]string{"ListSpCertificatesResponse", string(data)}, " ")
}
