package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNoteRequestBody 删除评论
type DeleteNoteRequestBody struct {

	// 删除评论的ID列表
	BatchIds []string `json:"batch_ids"`
}

func (o DeleteNoteRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNoteRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteNoteRequestBody", string(data)}, " ")
}
