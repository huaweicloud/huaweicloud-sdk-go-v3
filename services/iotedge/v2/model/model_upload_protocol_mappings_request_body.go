package model

import (
	"encoding/json"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
	"strings"
)

type UploadProtocolMappingsRequestBody struct {
	// 上传协议映射文件。当前仅支持xlsx/xls文件格式，且文件最大行数为10000行。
	File *def.FilePart `json:"file"`
}

func (o UploadProtocolMappingsRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UploadProtocolMappingsRequestBody struct{}"
	}

	return strings.Join([]string{"UploadProtocolMappingsRequestBody", string(data)}, " ")
}
