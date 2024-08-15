package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobScriptOrderListPage 脚本工单列表，分页数据
type JobScriptOrderListPage struct {

	// 总条数
	Total int64 `json:"total"`

	// 单页数据列表
	Data []JobScriptOrderListModel `json:"data"`
}

func (o JobScriptOrderListPage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobScriptOrderListPage struct{}"
	}

	return strings.Join([]string{"JobScriptOrderListPage", string(data)}, " ")
}
