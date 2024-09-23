package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PptPageInfo PPT图片元数据。
type PptPageInfo struct {

	// **参数解释**： 页面编号。 **约束限制**： 不涉及
	PageNo *int32 `json:"page_no,omitempty"`

	// **参数解释**： 页面对应图片文件ID。 **约束限制**： 不涉及 **取值范围**： 字符长度1-64位 **默认取值**： 不涉及
	FileId *string `json:"file_id,omitempty"`

	// **参数解释**： 页面对应图片文件ID **约束限制**： 不涉及 **取值范围**： 字符长度0-2048位 **默认取值**： 不涉及。
	PageContent *string `json:"page_content,omitempty"`
}

func (o PptPageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PptPageInfo struct{}"
	}

	return strings.Join([]string{"PptPageInfo", string(data)}, " ")
}
