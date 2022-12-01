package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 基础响应体。
type BaseResp struct {

	// 列表中的项目总数，与分页无关。
	TotalCount int32 `json:"total_count"`
}

func (o BaseResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseResp struct{}"
	}

	return strings.Join([]string{"BaseResp", string(data)}, " ")
}
