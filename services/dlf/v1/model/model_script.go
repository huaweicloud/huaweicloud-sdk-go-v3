package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Script struct {
	Name *string `json:"name,omitempty"`

	Path *string `json:"path,omitempty"`
}

func (o Script) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Script struct{}"
	}

	return strings.Join([]string{"Script", string(data)}, " ")
}
