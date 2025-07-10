package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandDesktopPoolReq 扩容桌面池请求。
type ExpandDesktopPoolReq struct {

	// 扩容桌面池的大小。
	Size int32 `json:"size"`
}

func (o ExpandDesktopPoolReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandDesktopPoolReq struct{}"
	}

	return strings.Join([]string{"ExpandDesktopPoolReq", string(data)}, " ")
}
