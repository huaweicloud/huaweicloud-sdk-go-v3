package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTsviReq 批量服务器更新虚拟会话IP配置请求内容
type UpdateTsviReq struct {

	// 批量请求列表，一次请求数量区间 [1, 20]
	Items []UpdateTsvi `json:"items"`
}

func (o UpdateTsviReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTsviReq struct{}"
	}

	return strings.Join([]string{"UpdateTsviReq", string(data)}, " ")
}
