package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type KeyAuthAuthorizerConfiguration struct {
	ApiKeys []ApiKeyInfo `json:"api_keys"`
}

func (o KeyAuthAuthorizerConfiguration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeyAuthAuthorizerConfiguration struct{}"
	}

	return strings.Join([]string{"KeyAuthAuthorizerConfiguration", string(data)}, " ")
}
