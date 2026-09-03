package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommonDto struct {

	// id
	Id *string `json:"id,omitempty"`

	// id对应的名称
	Name *string `json:"name,omitempty"`
}

func (o CommonDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommonDto struct{}"
	}

	return strings.Join([]string{"CommonDto", string(data)}, " ")
}
