package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublicScriptListPage 公共脚本列表分页数据
type PublicScriptListPage struct {

	// 总条数
	Total int64 `json:"total"`

	// 单页数据列表
	Data []PublicScriptListModel `json:"data"`
}

func (o PublicScriptListPage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicScriptListPage struct{}"
	}

	return strings.Join([]string{"PublicScriptListPage", string(data)}, " ")
}
