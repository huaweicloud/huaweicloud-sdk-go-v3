package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListEdgeNodeCertsResponse struct {

	// 节点上已关联的应用证书和设备证书的数目
	Count *int32 `json:"count,omitempty" xml:"count"`

	// 节点上的证书列表
	NodeCerts      *[]NodeCert `json:"node_certs,omitempty" xml:"node_certs"`
	HttpStatusCode int         `json:"-"`
}

func (o ListEdgeNodeCertsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEdgeNodeCertsResponse struct{}"
	}

	return strings.Join([]string{"ListEdgeNodeCertsResponse", string(data)}, " ")
}
