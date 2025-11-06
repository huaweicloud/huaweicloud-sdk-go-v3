package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCertsDetailRequest Request Object
type ShowCertsDetailRequest struct {

	// 指定待查询的集群ID。
	ClusterId string `json:"cluster_id"`

	// 指定查询的证书文件ID。
	CertId string `json:"cert_id"`
}

func (o ShowCertsDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCertsDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowCertsDetailRequest", string(data)}, " ")
}
