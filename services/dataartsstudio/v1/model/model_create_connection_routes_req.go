package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConnectionRoutesReq 创建路由的请求body体。
type CreateConnectionRoutesReq struct {

	// 路由名称，长度限制：1-64个字符。
	Name string `json:"name"`

	// 路由网段范围。
	Cidr string `json:"cidr"`
}

func (o CreateConnectionRoutesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectionRoutesReq struct{}"
	}

	return strings.Join([]string{"CreateConnectionRoutesReq", string(data)}, " ")
}
