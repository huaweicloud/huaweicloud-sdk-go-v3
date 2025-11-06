package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCertsRequest Request Object
type DeleteCertsRequest struct {

	// 指定删除证书文件的集群ID。
	ClusterId string `json:"cluster_id"`

	// 指定删除的证书文件ID。
	CertId string `json:"cert_id"`
}

func (o DeleteCertsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCertsRequest struct{}"
	}

	return strings.Join([]string{"DeleteCertsRequest", string(data)}, " ")
}
