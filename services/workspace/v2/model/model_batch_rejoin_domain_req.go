package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRejoinDomainReq 批量重新加域请求。
type BatchRejoinDomainReq struct {

	// 桌面ID列表。
	DesktopIds []string `json:"desktop_ids"`
}

func (o BatchRejoinDomainReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRejoinDomainReq struct{}"
	}

	return strings.Join([]string{"BatchRejoinDomainReq", string(data)}, " ")
}
