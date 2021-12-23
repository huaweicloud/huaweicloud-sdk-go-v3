package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 批量设置replaceDefiner请求体
type BatchReplaceDefinerReq struct {
	// 批量设置replaceDefiner请求列表

	Jobs []ReplaceDefinerInfo `json:"jobs"`
}

func (o BatchReplaceDefinerReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchReplaceDefinerReq struct{}"
	}

	return strings.Join([]string{"BatchReplaceDefinerReq", string(data)}, " ")
}
