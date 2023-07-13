package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadFromObsReq 从OBS导入元数据请求体
type UploadFromObsReq struct {

	// 元数据存储地址。
	MetadataPath string `json:"metadata_path"`

	// 元数据的名字。
	Name string `json:"name"`

	// 对元数据的描述。
	Description *string `json:"description,omitempty"`

	Encryption *UploadFromObsReqEncryption `json:"encryption,omitempty"`
}

func (o UploadFromObsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadFromObsReq struct{}"
	}

	return strings.Join([]string{"UploadFromObsReq", string(data)}, " ")
}
