package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type CreateMetadataReq struct {
	// 元数据存储地址。

	MetadataPath string `json:"metadataPath"`
	// 元数据的名字，限制为1-64个字符，且只能包含字母，数字或下划线。

	Name string `json:"name"`
	// 对元数据的描述。

	Description string `json:"description"`
	// 是否覆盖文件。

	IsOverwrite bool `json:"isOverwrite"`
	// 存储metadata的消息信息的对象。

	GesMetadata *interface{} `json:"gesMetadata"`
}

func (o CreateMetadataReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMetadataReq struct{}"
	}

	return strings.Join([]string{"CreateMetadataReq", string(data)}, " ")
}
