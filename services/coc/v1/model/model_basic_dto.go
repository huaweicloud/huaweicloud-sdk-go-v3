package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BasicDto 基础请求对象
type BasicDto struct {

	// id
	Id *string `json:"id,omitempty"`

	// code
	Code *string `json:"code,omitempty"`
}

func (o BasicDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BasicDto struct{}"
	}

	return strings.Join([]string{"BasicDto", string(data)}, " ")
}
