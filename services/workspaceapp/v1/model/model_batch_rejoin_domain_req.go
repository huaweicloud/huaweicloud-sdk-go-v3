package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRejoinDomainReq 批量重新加域请求。
type BatchRejoinDomainReq struct {

	// 批量唯一标识请求列表，一次请求数量区间 [1, 20]
	Items []string `json:"items"`
}

func (o BatchRejoinDomainReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRejoinDomainReq struct{}"
	}

	return strings.Join([]string{"BatchRejoinDomainReq", string(data)}, " ")
}
