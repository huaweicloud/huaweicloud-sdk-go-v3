package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ViewRow struct {

	// 视图行，包含多个视图，展示的时候根据实际的长度适配。
	ViewList *[]ViewBase `json:"view_list,omitempty"`

	// 标题。
	Title *string `json:"title,omitempty"`
}

func (o ViewRow) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ViewRow struct{}"
	}

	return strings.Join([]string{"ViewRow", string(data)}, " ")
}
