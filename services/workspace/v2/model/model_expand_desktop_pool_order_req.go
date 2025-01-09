package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandDesktopPoolOrderReq 扩容桌面池请求。
type ExpandDesktopPoolOrderReq struct {

	// 扩容桌面池的大小。
	Size *int32 `json:"size,omitempty"`

	// 要扩容的桌面池ID。
	PoolId string `json:"pool_id"`
}

func (o ExpandDesktopPoolOrderReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandDesktopPoolOrderReq struct{}"
	}

	return strings.Join([]string{"ExpandDesktopPoolOrderReq", string(data)}, " ")
}
