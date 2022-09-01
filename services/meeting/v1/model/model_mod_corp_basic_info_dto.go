package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 企业注册信息
type ModCorpBasicInfoDto struct {

	// 企业所在地,最大长度为255个字符。 maxLength：255。
	Address *string `json:"address,omitempty" xml:"address"`

	// 企业自动开户开关
	AutoUserCreate *bool `json:"autoUserCreate,omitempty" xml:"autoUserCreate"`
}

func (o ModCorpBasicInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModCorpBasicInfoDto struct{}"
	}

	return strings.Join([]string{"ModCorpBasicInfoDto", string(data)}, " ")
}
