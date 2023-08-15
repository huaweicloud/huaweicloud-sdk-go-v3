package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TokenProjectResult
type TokenProjectResult struct {

	// 项目名。
	Name string `json:"name"`

	// 项目ID。
	Id string `json:"id"`

	Domain *TokenProjectDomainResult `json:"domain"`
}

func (o TokenProjectResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TokenProjectResult struct{}"
	}

	return strings.Join([]string{"TokenProjectResult", string(data)}, " ")
}
