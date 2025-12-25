package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NotesDetailData 评论内容体
type NotesDetailData struct {

	// 评论的内容
	Content *string `json:"content,omitempty"`
}

func (o NotesDetailData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotesDetailData struct{}"
	}

	return strings.Join([]string{"NotesDetailData", string(data)}, " ")
}
