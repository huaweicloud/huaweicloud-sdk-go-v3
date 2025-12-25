package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IdInfo 分类信息
type IdInfo struct {

	// 分类id
	Id *string `json:"id,omitempty"`
}

func (o IdInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdInfo struct{}"
	}

	return strings.Join([]string{"IdInfo", string(data)}, " ")
}
