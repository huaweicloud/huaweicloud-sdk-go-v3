package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadKieResponseBodyMetadata 导出文件的描述信息
type DownloadKieResponseBodyMetadata struct {

	// 版本号
	Version *string `json:"version,omitempty"`

	// 导出文件的其他信息
	Annotations *interface{} `json:"annotations,omitempty"`
}

func (o DownloadKieResponseBodyMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadKieResponseBodyMetadata struct{}"
	}

	return strings.Join([]string{"DownloadKieResponseBodyMetadata", string(data)}, " ")
}
