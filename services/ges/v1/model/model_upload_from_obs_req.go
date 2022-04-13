package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UploadFromObsReq struct {
	// 元数据存储地址。

	MetadataPath string `json:"metadataPath"`
	// 元数据的名字。

	Name string `json:"name"`
	// 对元数据的描述。

	Description *string `json:"description,omitempty"`

	Encryption *EncryptionReq `json:"encryption,omitempty"`
}

func (o UploadFromObsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadFromObsReq struct{}"
	}

	return strings.Join([]string{"UploadFromObsReq", string(data)}, " ")
}
