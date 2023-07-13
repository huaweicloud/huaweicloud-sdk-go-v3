package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KnowledgePointInfo 习题知识点信息
type KnowledgePointInfo struct {

	// 知识点id
	Id string `json:"id"`

	// 知识点名称
	Name string `json:"name"`

	// 知识点顺序编号
	Sequence int32 `json:"sequence"`
}

func (o KnowledgePointInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KnowledgePointInfo struct{}"
	}

	return strings.Join([]string{"KnowledgePointInfo", string(data)}, " ")
}
