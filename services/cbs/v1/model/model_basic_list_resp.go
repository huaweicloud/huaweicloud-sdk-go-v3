package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BasicListResp
type BasicListResp struct {

	// 配额
	Quota int32 `json:"quota"`

	// 总数
	Total int32 `json:"total"`

	// 偏移
	Offset int32 `json:"offset"`

	// 返回数量
	Count int32 `json:"count"`
}

func (o BasicListResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BasicListResp struct{}"
	}

	return strings.Join([]string{"BasicListResp", string(data)}, " ")
}
