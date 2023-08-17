package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FilesCreateReq 创建文件请求。
type FilesCreateReq struct {

	// 文件名，不区分大小写，最大长度256，最小长度1。
	FileName string `json:"file_name"`

	// 文件内容MD5值，MD5值需要进行Base64编码。
	FileMd5 string `json:"file_md5"`

	// 文件总的大小，最小1，最大5368709120。
	FileSize int64 `json:"file_size"`

	// 文件类型（默认提取文件后缀）。
	FileType string `json:"file_type"`

	// 资产ID。
	AssetId string `json:"asset_id"`

	// 文件在资产中的分类。每种资产类型包含的文件分类不同。 * MAIN：主文件 * COVER：封面文件 * PAGE：内容页图片 * SAMPLE：样例音频 * OTHER：其他文件 * WHOLE_MODEL：全模型 * USER_MODIFIED_MODEL：用户上传模型 > * 资产类型为SCENE、ANIMATION、VIDEO、IMAGE、MATERIAL时，包含MAIN、COVER和OTHER > * 资产类型为PPT时，包含MAIN、COVER、PAGE和OTHER > * 资产类型为HUMAN_MODEL时，包含MAIN、COVER、WHOLE_MODEL、USER_MODIFIED_MODEL和OTHER > * 资产类型为VOICE_MODEL时，包含MAIN、SAMPLE(样例音频文件)和OTHER > * 资产类型为HUMAN_MODEL_2D时，包含MAIN、COVER、SAMPLE(动作样例)和OTHER(遮罩文件) > * 资产类型为BUSINESS_CARD_TEMPLET时，包含MAIN和COVER(名片效果图)
	AssetFileCategory string `json:"asset_file_category"`
}

func (o FilesCreateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FilesCreateReq struct{}"
	}

	return strings.Join([]string{"FilesCreateReq", string(data)}, " ")
}
