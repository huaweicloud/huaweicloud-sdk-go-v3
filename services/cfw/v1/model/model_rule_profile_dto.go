package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleProfileDto struct {

	// 域名url
	Url *string `json:"url,omitempty"`
}

func (o RuleProfileDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleProfileDto struct{}"
	}

	return strings.Join([]string{"RuleProfileDto", string(data)}, " ")
}
