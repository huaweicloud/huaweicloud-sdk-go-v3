package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModCorpBasicInfoDto 企业注册信息
type ModCorpBasicInfoDto struct {

	// 企业所在地，最大长度为255个字符。
	Address *string `json:"address,omitempty"`

	// 企业自动开户开关。
	AutoUserCreate *bool `json:"autoUserCreate,omitempty"`
}

func (o ModCorpBasicInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModCorpBasicInfoDto struct{}"
	}

	return strings.Join([]string{"ModCorpBasicInfoDto", string(data)}, " ")
}
