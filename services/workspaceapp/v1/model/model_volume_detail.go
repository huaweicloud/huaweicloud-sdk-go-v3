package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VolumeDetail 磁盘信息。
type VolumeDetail struct {

	// 标识磁盘是否加密，如果为1就是加密，0非加密。
	EncryptFlag *string `json:"encrypt_flag,omitempty"`

	// 如果磁盘加密，传递的密钥。
	KmsKey *string `json:"kms_key,omitempty"`

	// 如果磁盘加密，传递的密钥。
	KeyAlias *string `json:"key_alias,omitempty"`

	Type *VolumeType `json:"type,omitempty"`

	// 磁盘容量，单位GB。
	Size *int32 `json:"size,omitempty"`

	// 如果磁盘加密，授权ID。
	KmsGrantId *string `json:"kms_grant_id,omitempty"`

	// 挂载目录。
	Device *string `json:"device,omitempty"`

	// 磁盘表唯一标识ID。
	Id *string `json:"id,omitempty"`

	// 磁盘ID。
	VolumeId *string `json:"volume_id,omitempty"`

	// 专属分布式存储池id。
	ClusterId *string `json:"cluster_id,omitempty"`
}

func (o VolumeDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeDetail struct{}"
	}

	return strings.Join([]string{"VolumeDetail", string(data)}, " ")
}
