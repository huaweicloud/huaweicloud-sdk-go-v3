package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PageResp API列表查询响应基类。
type PageResp struct {

	// 总数。
	Count *int32 `json:"count,omitempty"`
}

func (o PageResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageResp struct{}"
	}

	return strings.Join([]string{"PageResp", string(data)}, " ")
}
