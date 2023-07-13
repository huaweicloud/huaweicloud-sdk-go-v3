package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchTaskData 分子搜索任务的请求体
type SearchTaskData struct {

	// 分子SMILES表达式
	Smiles string `json:"smiles"`

	// 搜索使用到的数据库集合
	Databases []string `json:"databases"`

	// 期望最大返回条目数（排序后取Top）
	TopN *int32 `json:"top_n,omitempty"`
}

func (o SearchTaskData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchTaskData struct{}"
	}

	return strings.Join([]string{"SearchTaskData", string(data)}, " ")
}
