package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SearchQaPairsReq struct {

	// 主题名称
	Domain *string `json:"domain,omitempty"`

	// 返回匹配度最高的数据条数
	Top *int32 `json:"top,omitempty"`
}

func (o SearchQaPairsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchQaPairsReq struct{}"
	}

	return strings.Join([]string{"SearchQaPairsReq", string(data)}, " ")
}
