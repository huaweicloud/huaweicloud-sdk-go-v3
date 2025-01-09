package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConfigInfoReq 批量查询企业配置请求。
type ListConfigInfoReq struct {

	// 查询企业配置请求的key，一次请求数量区间 [0, 100]。
	Items []string `json:"items"`
}

func (o ListConfigInfoReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigInfoReq struct{}"
	}

	return strings.Join([]string{"ListConfigInfoReq", string(data)}, " ")
}
