package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyColumnInfoReq 修改列信息请求体
type ModifyColumnInfoReq struct {

	// 列信息
	ColumnMapInfos []ColumnMappingInfo `json:"column_map_infos"`
}

func (o ModifyColumnInfoReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyColumnInfoReq struct{}"
	}

	return strings.Join([]string{"ModifyColumnInfoReq", string(data)}, " ")
}
