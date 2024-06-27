package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BasePageInfoReview struct {
	Total *int64 `json:"total,omitempty"`

	List *[]Review `json:"list,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Pages *int32 `json:"pages,omitempty"`

	Size *int32 `json:"size,omitempty"`
}

func (o BasePageInfoReview) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BasePageInfoReview struct{}"
	}

	return strings.Join([]string{"BasePageInfoReview", string(data)}, " ")
}
