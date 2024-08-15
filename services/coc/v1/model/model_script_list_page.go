package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScriptListPage 脚本列表分页数据
type ScriptListPage struct {

	// 总条数
	Total int64 `json:"total"`

	// 单页数据列表
	Data []ScriptListModel `json:"data"`
}

func (o ScriptListPage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScriptListPage struct{}"
	}

	return strings.Join([]string{"ScriptListPage", string(data)}, " ")
}
