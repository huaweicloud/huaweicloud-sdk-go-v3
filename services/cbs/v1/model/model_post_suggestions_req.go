package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type PostSuggestionsReq struct {
	// 用户输入的问题，长度为1~512。

	Question string `json:"question"`
	// 最多提示条数，默认为5，取值范围[1,10]。

	Top *int32 `json:"top,omitempty"`
}

func (o PostSuggestionsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostSuggestionsReq struct{}"
	}

	return strings.Join([]string{"PostSuggestionsReq", string(data)}, " ")
}
