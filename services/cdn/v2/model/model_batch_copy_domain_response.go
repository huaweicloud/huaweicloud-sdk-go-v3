package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCopyDomainResponse Response Object
type BatchCopyDomainResponse struct {

	// 成功响应体
	Result         *[]BatchCopyResultVo `json:"result,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o BatchCopyDomainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCopyDomainResponse struct{}"
	}

	return strings.Join([]string{"BatchCopyDomainResponse", string(data)}, " ")
}
