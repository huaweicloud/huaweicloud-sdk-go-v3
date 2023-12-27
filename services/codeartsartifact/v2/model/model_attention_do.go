package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AttentionDo struct {

	// 关注/取消关注
	Attention string `json:"attention"`

	// 组件id列表
	Ids []string `json:"ids"`

	// 格式
	Format *string `json:"format,omitempty"`
}

func (o AttentionDo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttentionDo struct{}"
	}

	return strings.Join([]string{"AttentionDo", string(data)}, " ")
}
