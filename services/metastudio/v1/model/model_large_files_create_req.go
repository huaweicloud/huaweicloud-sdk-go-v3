package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LargeFilesCreateReq 创建大文件请求。
type LargeFilesCreateReq struct {

	// 文件名，不区分大小写，最大长度256，最小长度1。
	FileName string `json:"file_name"`

	// 文件总的大小，最小1，最大536870912000。
	FileSize *int64 `json:"file_size,omitempty"`

	// 文件类型（默认提取文件后缀）。
	FileType string `json:"file_type"`

	// 资产ID。
	AssetId string `json:"asset_id"`

	// 文件在资产中的分类。每种资产类型包含的文件分类不同。 * MAIN：主文件 * OTHER：其他文件
	AssetFileCategory string `json:"asset_file_category"`

	// ORI4K文件分段上传数量，默认值为1
	FileMultipartCount *int32 `json:"file_multipart_count,omitempty"`
}

func (o LargeFilesCreateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LargeFilesCreateReq struct{}"
	}

	return strings.Join([]string{"LargeFilesCreateReq", string(data)}, " ")
}
