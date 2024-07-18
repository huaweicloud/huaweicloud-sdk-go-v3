package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateServerRequestServerCertificate 服务端证书
type UpdateServerRequestServerCertificate struct {

	// 服务端证书ID,为CCM服务中的证书ID
	Id *string `json:"id,omitempty"`
}

func (o UpdateServerRequestServerCertificate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerRequestServerCertificate struct{}"
	}

	return strings.Join([]string{"UpdateServerRequestServerCertificate", string(data)}, " ")
}
