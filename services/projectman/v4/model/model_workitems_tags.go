package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WorkitemsTags struct {

	// 标签id
	Id *string `json:"id,omitempty" xml:"id"`

	// 标签名
	Name *string `json:"name,omitempty" xml:"name"`
}

func (o WorkitemsTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkitemsTags struct{}"
	}

	return strings.Join([]string{"WorkitemsTags", string(data)}, " ")
}
