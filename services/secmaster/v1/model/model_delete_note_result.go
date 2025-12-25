package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNoteResult 删除评论结果对象
type DeleteNoteResult struct {

	// 成功删除的 note ID 列表
	SuccessIds *[]string `json:"success_ids,omitempty"`

	// 删除失败的 note ID 列表
	ErrorIds *[]string `json:"error_ids,omitempty"`
}

func (o DeleteNoteResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNoteResult struct{}"
	}

	return strings.Join([]string{"DeleteNoteResult", string(data)}, " ")
}
