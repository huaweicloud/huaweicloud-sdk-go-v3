package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BasePageInfoTestCase struct {
	Limit *int32 `json:"limit,omitempty"`

	List *[]TestCase `json:"list,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Pages *int32 `json:"pages,omitempty"`

	Size *int32 `json:"size,omitempty"`

	Total *int64 `json:"total,omitempty"`
}

func (o BasePageInfoTestCase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BasePageInfoTestCase struct{}"
	}

	return strings.Join([]string{"BasePageInfoTestCase", string(data)}, " ")
}
