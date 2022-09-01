package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 日志格式
type AccessConfigFormatCreate struct {
	Single *AccessConfigFormatSingleCreate `json:"single,omitempty" xml:"single"`

	Multi *AccessConfigFormatMutilCreate `json:"multi,omitempty" xml:"multi"`
}

func (o AccessConfigFormatCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfigFormatCreate struct{}"
	}

	return strings.Join([]string{"AccessConfigFormatCreate", string(data)}, " ")
}
