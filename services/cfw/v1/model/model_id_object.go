package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IdObject struct {

	// id值
	Id *string `json:"id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`
}

func (o IdObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdObject struct{}"
	}

	return strings.Join([]string{"IdObject", string(data)}, " ")
}
