package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReleaseRepository struct {

	// 软件包id
	Id *string `json:"id,omitempty" xml:"id"`

	// 软件包名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 软件包下载地址
	DownloadPath *string `json:"download_path,omitempty" xml:"download_path"`

	// 软件包大小
	Size *string `json:"size,omitempty" xml:"size"`

	// 文件类型
	FileType *string `json:"file_type,omitempty" xml:"file_type"`

	// 创建时间
	Created *string `json:"created,omitempty" xml:"created"`

	// 修改时间
	Updated *string `json:"updated,omitempty" xml:"updated"`
}

func (o ReleaseRepository) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReleaseRepository struct{}"
	}

	return strings.Join([]string{"ReleaseRepository", string(data)}, " ")
}
