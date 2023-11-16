package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerIdSet 服务器ID列表集合。
type ServerIdSet struct {

	// 批量请求的服务器ID列表，一次请求数量区间 [1, 20]。
	Items []string `json:"items"`
}

func (o ServerIdSet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerIdSet struct{}"
	}

	return strings.Join([]string{"ServerIdSet", string(data)}, " ")
}
