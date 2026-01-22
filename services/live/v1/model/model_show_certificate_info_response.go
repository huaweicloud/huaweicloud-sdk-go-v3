package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCertificateInfoResponse Response Object
type ShowCertificateInfoResponse struct {

	// 查询结果的总数量
	Total float32 `json:"total,omitempty"`

	// 证书信息列表
	CertificatesInfo *[]CertInfoResp `json:"certificates_info,omitempty"`
	HttpStatusCode   int             `json:"-"`
}

func (o ShowCertificateInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCertificateInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowCertificateInfoResponse", string(data)}, " ")
}
