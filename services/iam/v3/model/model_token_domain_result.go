package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TokenDomainResult
type TokenDomainResult struct {

	// 用户所属账号名。
	Name string `json:"name"`

	// 用户所属账号ID。
	Id string `json:"id"`
}

func (o TokenDomainResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TokenDomainResult struct{}"
	}

	return strings.Join([]string{"TokenDomainResult", string(data)}, " ")
}
