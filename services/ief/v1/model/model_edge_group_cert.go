package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 边缘节点组证书
type EdgeGroupCert struct {

	// 证书ID
	Id *string `json:"id,omitempty"`

	// 证书名称
	Name *string `json:"name,omitempty"`

	// 证书描述
	Description *string `json:"description,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 证书绑定的边缘节点组ID
	GroupId *string `json:"group_id,omitempty"`

	// 证书是否处于删除中
	IsDeleted *bool `json:"is_deleted,omitempty"`

	// 证书所属账号的项目ID
	ProjectId *string `json:"project_id,omitempty"`

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
	Package *string `json:"package,omitempty"`

	// 证书有效期持续时间
	CertRemainingValidTime *int32 `json:"cert_remaining_valid_time,omitempty"`
}

func (o EdgeGroupCert) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EdgeGroupCert struct{}"
	}

	return strings.Join([]string{"EdgeGroupCert", string(data)}, " ")
}
