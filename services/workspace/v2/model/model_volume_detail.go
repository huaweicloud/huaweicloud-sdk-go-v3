package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VolumeDetail 磁盘信息。
type VolumeDetail struct {

	// 标识磁盘是否加密，如果为1就是加密。
	EncryptFlag *string `json:"encrypt_flag,omitempty"`

	// 如果磁盘加密，传递的密钥。
	KmsKey *string `json:"kms_key,omitempty"`

	// 如果磁盘加密，传递的密钥。
	KeyAlias *string `json:"key_alias,omitempty"`

	// 桌面数据盘对应的磁盘类型，需要与系统所提供的磁盘类型相匹配。  - SAS：高IO。 - SSD：超高IO。
	Type string `json:"type"`

	// 磁盘容量，单位GB。
	Size int32 `json:"size"`

	// 如果磁盘加密，授权ID。
	KmsGrantId *string `json:"kms_grant_id,omitempty"`

	// 挂载目录。
	Device *string `json:"device,omitempty"`

	// 磁盘表唯一标识ID。
	Id *string `json:"id,omitempty"`

	// 磁盘ID。
	VolumeId *string `json:"volume_id,omitempty"`

	// 磁盘计费资源ID。
	BillResourceId *string `json:"bill_resource_id,omitempty"`

	// 磁盘的创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 磁盘名
	DisplayName *string `json:"display_name,omitempty"`

	// 云服务器系统盘对应的存储池的ID。
	ClusterId *string `json:"cluster_id,omitempty"`

	// 规格
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`
}

func (o VolumeDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeDetail struct{}"
	}

	return strings.Join([]string{"VolumeDetail", string(data)}, " ")
}
