package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UploadFromObsReq struct {

	// 元数据存储地址。
	MetadataPath string `json:"metadataPath" xml:"metadataPath"`

	// 元数据的名字。
	Name string `json:"name" xml:"name"`

	// 对元数据的描述。
	Description *string `json:"description,omitempty" xml:"description"`

	Encryption *EncryptionReq `json:"encryption,omitempty" xml:"encryption"`
}

func (o UploadFromObsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadFromObsReq struct{}"
	}

	return strings.Join([]string{"UploadFromObsReq", string(data)}, " ")
}
