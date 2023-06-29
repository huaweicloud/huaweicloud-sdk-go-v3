package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssetFileInfo 文件信息。
type AssetFileInfo struct {

	// 文件ID。
	FileId *string `json:"file_id,omitempty"`

	// 文件名创建文件时候不区分大小写，最大长度256，最小长度1。
	FileName *string `json:"file_name,omitempty"`

	// 文件内容MD5值,固定32位。
	FileMd5 *string `json:"file_md5,omitempty"`

	// 文件总的大小，最小1，最大5368709120。
	FileSize *int64 `json:"file_size,omitempty"`

	// 文件类型（默认提取文件后缀）。
	FileType *string `json:"file_type,omitempty"`

	// 文件在资产中的分类。每种资产类型包含的文件分类不同。 * MAIN： 主文件 * COVER： 封面文件 * PAGE: 内容页图片 * SAMPLE：样例音频 * OTHER： 其他文件 > * 资产类型为SCENE、ANIMATION、VIDEO、IMAGE、MATERIAL时，包含MAIN、COVER和OTHER > * 资产类型为PPT时，包含MAIN、COVER、PAGE和OTHER > * 资产类型为HUMAN_MODEL时，包含MAIN、COVER和OTHER > * 资产类型为VOICE_MODEL时，包含MAIN、SAMPLE和OTHER
	AssetFileCategory *string `json:"asset_file_category,omitempty"`

	// 文件下载URL,有效期为24小时。
	DownloadUrl *string `json:"download_url,omitempty"`
}

func (o AssetFileInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssetFileInfo struct{}"
	}

	return strings.Join([]string{"AssetFileInfo", string(data)}, " ")
}
