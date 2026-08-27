package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModelPageResp 模型管理模块 API 列表查询响应基类。
type ModelPageResp struct {

	// 总数。
	Total *int32 `json:"total,omitempty"`
}

func (o ModelPageResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelPageResp struct{}"
	}

	return strings.Join([]string{"ModelPageResp", string(data)}, " ")
}
