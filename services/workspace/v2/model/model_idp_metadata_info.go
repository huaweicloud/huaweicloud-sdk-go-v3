package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IdpMetadataInfo IDP元数据信息。
type IdpMetadataInfo struct {

	// IDP元数据id。
	Id *string `json:"id,omitempty"`

	// IDP元数据文件名称。
	FileName *string `json:"file_name,omitempty"`

	// IDP元数据内容（base64）。
	Content *string `json:"content,omitempty"`

	// 元数据大小（字节数），最大1M。
	ContentLength float32 `json:"content_length,omitempty"`

	// 创建时间戳。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新时间戳。
	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o IdpMetadataInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdpMetadataInfo struct{}"
	}

	return strings.Join([]string{"IdpMetadataInfo", string(data)}, " ")
}
