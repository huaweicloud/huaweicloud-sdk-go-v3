package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Metadata 元数据
type Metadata struct {

	// 元数据 ID。
	Id string `json:"id"`

	// 元数据名称。
	Name string `json:"name"`

	// 元数据 描述。
	Description string `json:"description"`

	// 元数据是否可用。
	Status string `json:"status"`

	// 元数据对应路径。
	MetadataPath string `json:"metadataPath"`

	// 元数据创建时间戳。
	CreateTimestamp *string `json:"createTimestamp,omitempty"`

	// 元数据最后更新时间戳。
	LastUpdateTimestamp *string `json:"lastUpdateTimestamp,omitempty"`
}

func (o Metadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Metadata struct{}"
	}

	return strings.Join([]string{"Metadata", string(data)}, " ")
}
