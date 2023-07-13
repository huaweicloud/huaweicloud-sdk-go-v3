package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReleaseRepository struct {

	// 软件包id
	Id *string `json:"id,omitempty"`

	// 软件包名称
	Name *string `json:"name,omitempty"`

	// 软件包下载地址
	DownloadPath *string `json:"download_path,omitempty"`

	// 软件包大小
	Size *string `json:"size,omitempty"`

	// 软件包类型名称
	CategoryName *string `json:"category_name,omitempty"`

	// 文件类型
	FileType *string `json:"file_type,omitempty"`

	// 创建时间
	Created *string `json:"created,omitempty"`

	// 修改时间
	Updated *string `json:"updated,omitempty"`
}

func (o ReleaseRepository) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReleaseRepository struct{}"
	}

	return strings.Join([]string{"ReleaseRepository", string(data)}, " ")
}
