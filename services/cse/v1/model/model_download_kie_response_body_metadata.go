package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 导出文件的描述信息
type DownloadKieResponseBodyMetadata struct {

	// 版本号
	Version *string `json:"version,omitempty" xml:"version"`

	// 导出文件的其他信息
	Annotations *interface{} `json:"annotations,omitempty" xml:"annotations"`
}

func (o DownloadKieResponseBodyMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadKieResponseBodyMetadata struct{}"
	}

	return strings.Join([]string{"DownloadKieResponseBodyMetadata", string(data)}, " ")
}
