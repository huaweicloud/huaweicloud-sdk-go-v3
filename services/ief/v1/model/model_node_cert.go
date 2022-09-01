package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodeCert struct {

	// 证书id
	Id *string `json:"id,omitempty" xml:"id"`

	// 证书名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 证书的描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 证书的创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty" xml:"created_at"`

	// 节点id
	NodeId *string `json:"node_id,omitempty" xml:"node_id"`

	// 证书类型，包含： - system：创建节点时会默认创建一套系统证书 - application：应用证书 - device：设备证书
	Type *string `json:"type,omitempty" xml:"type"`

	// 证书序列号
	SerialNum *string `json:"serial_num,omitempty" xml:"serial_num"`
}

func (o NodeCert) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeCert struct{}"
	}

	return strings.Join([]string{"NodeCert", string(data)}, " ")
}
