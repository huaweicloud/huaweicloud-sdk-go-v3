package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PptPageInfo PPT图片元数据。
type PptPageInfo struct {

	// 页面编号。
	PageNo *int32 `json:"page_no,omitempty"`

	// 页面对应图片文件ID。
	FileId *string `json:"file_id,omitempty"`

	// 讲解词（从备注中提取）。
	PageContent *string `json:"page_content,omitempty"`
}

func (o PptPageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PptPageInfo struct{}"
	}

	return strings.Join([]string{"PptPageInfo", string(data)}, " ")
}
