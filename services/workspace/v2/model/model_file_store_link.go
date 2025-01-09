package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FileStoreLink 文件存储路径支持OBS或者外部链接，store_type指定实际生效的配置,bucket_store和file_link必须设置其一。
type FileStoreLink struct {
	StoreType *FileStoreTypeEnum `json:"store_type"`

	BucketStore *BucketStore `json:"bucket_store,omitempty"`

	// 文件下载完整路径。
	FileLink *string `json:"file_link,omitempty"`
}

func (o FileStoreLink) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileStoreLink struct{}"
	}

	return strings.Join([]string{"FileStoreLink", string(data)}, " ")
}
