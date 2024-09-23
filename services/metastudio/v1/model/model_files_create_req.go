package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FilesCreateReq 创建文件请求。
type FilesCreateReq struct {

	// **参数解释**： 文件名。 **约束限制**： 不区分大小写。 **取值范围**： 字符长度1-256位。 **默认取值**： 不涉及
	FileName string `json:"file_name"`

	// **参数解释**： 文件内容MD5值。按照RFC 1864标准计算出消息体的MD5摘要字符串，即消息体128-bit MD5值经过base64编码后得到的字符串。 **约束限制**： 不涉及 **取值范围**： 字符长度24位。 **默认取值**： 不涉及
	FileMd5 string `json:"file_md5"`

	// **参数解释**： 文件总的大小。 **约束限制**： 最大支持5GB  **默认取值**： 不涉及
	FileSize int64 `json:"file_size"`

	// **参数解释**： 文件类型 **约束限制**： 不涉及 **取值范围**： 字符长度1-64位。 **默认取值**： 默认提取文件后缀。
	FileType string `json:"file_type"`

	// **参数解释**： 本平台资产ID。 **约束限制**： 不涉及 **取值范围**： 字符长度1-64位。 **默认取值**： 不涉及
	AssetId string `json:"asset_id"`

	// **参数解释**： 文件在资产中的分类。每种资产类型包含的文件分类不同。 * MAIN：主文件 * COVER：封面文件 * PAGE：PPT内容页图片文件 * SAMPLE：样例音频或样例动作文件 * OTHER：其他文件 * WHOLE_MODEL：全模型文件（3D数字人） * USER_MODIFIED_MODEL：用户上传模型（3D数字人） > * 资产类型为SCENE、ANIMATION、VIDEO、IMAGE、MATERIAL时，包含MAIN、COVER和OTHER > * 资产类型为PPT时，包含MAIN、COVER、PAGE和OTHER > * 资产类型为HUMAN_MODEL时，包含MAIN、COVER、WHOLE_MODEL、USER_MODIFIED_MODEL和OTHER > * 资产类型为VOICE_MODEL时，包含MAIN、SAMPLE(样例音频文件)和OTHER > * 资产类型为HUMAN_MODEL_2D时，包含MAIN、COVER、SAMPLE(动作样例)和OTHER > * 资产类型为BUSINESS_CARD_TEMPLET时，包含MAIN和COVER(名片效果图) > * 资产类型为IMAGE时，包含MAIN > * 资产类型为VIDEO时，包含MAIN、COVER  **约束限制**： 一个资产中MAIN文件只有一个，且必须有一个 **取值范围**： 字符长度1-128位。 **默认取值**： 不涉及
	AssetFileCategory string `json:"asset_file_category"`
}

func (o FilesCreateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FilesCreateReq struct{}"
	}

	return strings.Join([]string{"FilesCreateReq", string(data)}, " ")
}
