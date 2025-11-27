package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListMeta struct {

	// 服务器内部版本标识符
	ResourceVersion *string `json:"resourceVersion,omitempty"`

	// 表示当前请求返回的结果是否还有下一页数据
	Continue *string `json:"continue,omitempty"`

	// 表示在当前响应之后，列表中还有多少项未被包含进来
	RemainingItemCount *string `json:"remainingItemCount,omitempty"`
}

func (o ListMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMeta struct{}"
	}

	return strings.Join([]string{"ListMeta", string(data)}, " ")
}
