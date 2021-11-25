package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApiGroupInfoPage struct {
	// 本次返回的列表长度

	Size int32 `json:"size"`
	// 满足条件的记录数

	Total int64 `json:"total"`
	// 分组列表

	Groups *[]ApiGroupInfoPage `json:"groups,omitempty"`
}

func (o ApiGroupInfoPage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiGroupInfoPage struct{}"
	}

	return strings.Join([]string{"ApiGroupInfoPage", string(data)}, " ")
}
