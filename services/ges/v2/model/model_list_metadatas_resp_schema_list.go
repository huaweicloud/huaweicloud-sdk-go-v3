package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListMetadatasRespSchemaList struct {

	// 元数据 ID。
	Id *string `json:"id,omitempty"`

	// 元数据名称。
	Name *string `json:"name,omitempty"`

	// 元数据创建时间
	StartTime *string `json:"start_time,omitempty"`

	// 元数据最后更新时间
	LastUpdateTime *string `json:"last_update_time,omitempty"`

	// 元数据是否加密
	Encrypted *bool `json:"encrypted,omitempty"`

	// 秘钥名称
	MasterKeyName *string `json:"master_key_name,omitempty"`

	// 秘钥id
	MasterKeyId *string `json:"master_key_id,omitempty"`

	// 元数据 描述。
	Description *string `json:"description,omitempty"`

	// 元数据对应路径。
	MetadataPath *string `json:"metadata_path,omitempty"`

	// 元数据是否可用。
	Status *string `json:"status,omitempty"`
}

func (o ListMetadatasRespSchemaList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetadatasRespSchemaList struct{}"
	}

	return strings.Join([]string{"ListMetadatasRespSchemaList", string(data)}, " ")
}
