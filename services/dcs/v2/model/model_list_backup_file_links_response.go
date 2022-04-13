package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListBackupFileLinksResponse struct {
	// OBS桶内文件路径。

	FilePath *string `json:"file_path,omitempty"`
	// OBS桶名。

	BucketName *string `json:"bucket_name,omitempty"`
	// 备份文件下链接集合，链接数最大为64个。

	Links          *[]LinksItem `json:"links,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListBackupFileLinksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackupFileLinksResponse struct{}"
	}

	return strings.Join([]string{"ListBackupFileLinksResponse", string(data)}, " ")
}
