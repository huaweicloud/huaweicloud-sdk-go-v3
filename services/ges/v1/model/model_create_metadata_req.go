package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type CreateMetadataReq struct {

	// 元数据存储地址。
	MetadataPath string `json:"metadataPath" xml:"metadataPath"`

	// 元数据的名字，限制为1-64个字符，且只能包含字母，数字或下划线。
	Name string `json:"name" xml:"name"`

	// 对元数据的描述。
	Description string `json:"description" xml:"description"`

	// 是否覆盖文件。
	IsOverwrite bool `json:"isOverwrite" xml:"isOverwrite"`

	// 存储metadata的消息信息的对象。
	GesMetadata *interface{} `json:"gesMetadata" xml:"gesMetadata"`
}

func (o CreateMetadataReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMetadataReq struct{}"
	}

	return strings.Join([]string{"CreateMetadataReq", string(data)}, " ")
}
