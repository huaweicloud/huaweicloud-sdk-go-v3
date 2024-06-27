package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BasePageInfoTestPoint struct {
	Limit *int32 `json:"limit,omitempty"`

	List *[]TestPoint `json:"list,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Pages *int32 `json:"pages,omitempty"`

	Size *int32 `json:"size,omitempty"`

	Total *int64 `json:"total,omitempty"`
}

func (o BasePageInfoTestPoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BasePageInfoTestPoint struct{}"
	}

	return strings.Join([]string{"BasePageInfoTestPoint", string(data)}, " ")
}
