package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccessConfigFormatCreate 日志格式
type AccessConfigFormatCreate struct {
	Single *AccessConfigFormatSingleCreate `json:"single,omitempty"`

	Multi *AccessConfigFormatMutilCreate `json:"multi,omitempty"`
}

func (o AccessConfigFormatCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfigFormatCreate struct{}"
	}

	return strings.Join([]string{"AccessConfigFormatCreate", string(data)}, " ")
}
