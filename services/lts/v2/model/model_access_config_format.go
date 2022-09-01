package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 日志格式
type AccessConfigFormat struct {
	Single *AccessConfigFormatSingle `json:"single,omitempty" xml:"single"`

	Multi *AccessConfigFormatMutil `json:"multi,omitempty" xml:"multi"`
}

func (o AccessConfigFormat) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfigFormat struct{}"
	}

	return strings.Join([]string{"AccessConfigFormat", string(data)}, " ")
}
