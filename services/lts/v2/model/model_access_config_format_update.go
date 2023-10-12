package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AccessConfigFormatUpdate struct {
	Single *AccessConfigFormatSingle `json:"single,omitempty"`

	Multi *AccessConfigFormatMutil `json:"multi,omitempty"`
}

func (o AccessConfigFormatUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfigFormatUpdate struct{}"
	}

	return strings.Join([]string{"AccessConfigFormatUpdate", string(data)}, " ")
}
