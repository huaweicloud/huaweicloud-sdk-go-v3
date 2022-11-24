package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateEdgeNodeCertsResponse struct {

	// 证书id
	Id *string `json:"id,omitempty"`

	// 证书名称
	Name *string `json:"name,omitempty"`

	// 证书的描述
	Description *string `json:"description,omitempty"`

	// 证书的创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 节点id
	NodeId *string `json:"node_id,omitempty"`

	// 证书类型，包含： - system：创建节点时会默认创建一套系统证书 - application：应用证书 - device：设备证书
	Type *string `json:"type,omitempty"`

	// 证书序列号
	SerialNum *string `json:"serial_num,omitempty"`

	// 根证书
	Ca *string `json:"ca,omitempty"`

	// 证书
	Certificate *string `json:"certificate,omitempty"`

	// 私钥
	PrivateKey *string `json:"private_key,omitempty"`

	// 将证书文件certificate/ca/private_key打成.tar.gz包后用base64编码的字符串。 使用时请使用base64解码成.tar.gz包。
	Package        *string `json:"package,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateEdgeNodeCertsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEdgeNodeCertsResponse struct{}"
	}

	return strings.Join([]string{"CreateEdgeNodeCertsResponse", string(data)}, " ")
}
