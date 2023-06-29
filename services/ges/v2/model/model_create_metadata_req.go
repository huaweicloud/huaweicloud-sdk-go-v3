package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMetadataReq 创建元数据请求体
type CreateMetadataReq struct {

	// 元数据存储地址。
	MetadataPath string `json:"metadata_path"`

	// 元数据的名字，限制为1-64个字符，且只能包含字母，数字或下划线。
	Name string `json:"name"`

	// 对元数据的描述。
	Description string `json:"description"`

	// 是否覆盖文件。
	IsOverwrite bool `json:"is_overwrite"`

	GesMetadata *CreateMetadataReqGesMetadata `json:"ges_metadata"`
}

func (o CreateMetadataReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMetadataReq struct{}"
	}

	return strings.Join([]string{"CreateMetadataReq", string(data)}, " ")
}
