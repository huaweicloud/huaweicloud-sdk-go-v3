package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCheckRejoinDomainReq 批量检查加域请求。
type BatchCheckRejoinDomainReq struct {

	// 桌面ID列表。
	DesktopIds []string `json:"desktop_ids"`
}

func (o BatchCheckRejoinDomainReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCheckRejoinDomainReq struct{}"
	}

	return strings.Join([]string{"BatchCheckRejoinDomainReq", string(data)}, " ")
}
