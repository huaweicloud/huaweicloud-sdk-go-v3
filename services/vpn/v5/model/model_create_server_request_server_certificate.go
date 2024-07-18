package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateServerRequestServerCertificate 服务端证书。 隧道协议类型是SSL时，必填
type CreateServerRequestServerCertificate struct {

	// 服务端证书ID,为CCM服务中的证书ID
	Id string `json:"id"`
}

func (o CreateServerRequestServerCertificate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServerRequestServerCertificate struct{}"
	}

	return strings.Join([]string{"CreateServerRequestServerCertificate", string(data)}, " ")
}
