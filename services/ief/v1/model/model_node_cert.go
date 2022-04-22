package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodeCert struct {

	// 证书id
	Id *string `json:"id,omitempty"`

	// 证书名称
	Name *string `json:"name,omitempty"`

	// 证书的描述
	Description *string `json:"description,omitempty"`

	// 证书的创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 节点id
	NodeId *string `json:"node_id,omitempty"`

	// 证书类型，包含： - system：创建节点时会默认创建一套系统证书； - application：应用证书； - device：设备证书；
	Type *string `json:"type,omitempty"`

	// 证书序列号
	SerialNum *string `json:"serial_num,omitempty"`
}

func (o NodeCert) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeCert struct{}"
	}

	return strings.Join([]string{"NodeCert", string(data)}, " ")
}
