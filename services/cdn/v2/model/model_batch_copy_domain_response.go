package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCopyDomainResponse Response Object
type BatchCopyDomainResponse struct {

	// 复制配置结果。
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
