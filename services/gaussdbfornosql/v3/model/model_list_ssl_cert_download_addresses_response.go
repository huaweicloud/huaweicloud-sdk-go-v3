package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSslCertDownloadAddressesResponse Response Object
type ListSslCertDownloadAddressesResponse struct {

	// **参数解释：** 证书列表。 **取值范围：** 不涉及。
	Certs          *[]CertInfo `json:"certs,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListSslCertDownloadAddressesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSslCertDownloadAddressesResponse struct{}"
	}

	return strings.Join([]string{"ListSslCertDownloadAddressesResponse", string(data)}, " ")
}
