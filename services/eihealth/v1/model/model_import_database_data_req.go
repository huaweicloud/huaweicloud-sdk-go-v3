package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 导入数据请求体
type ImportDatabaseDataReq struct {

	// 导入文件l路径列表
	Files []string `json:"files"`

	// 分隔符，常见分隔符为, ;
	Delimiter string `json:"delimiter"`

	// 跳过的header行数
	SkipLines int32 `json:"skip_lines"`
}

func (o ImportDatabaseDataReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportDatabaseDataReq struct{}"
	}

	return strings.Join([]string{"ImportDatabaseDataReq", string(data)}, " ")
}
