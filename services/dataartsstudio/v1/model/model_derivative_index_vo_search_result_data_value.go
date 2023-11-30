package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DerivativeIndexVoSearchResultDataValue struct {

	// DerivativeIndexVO数组
	Records *[]DerivativeIndexVo `json:"records,omitempty"`

	// 总数
	Total *int32 `json:"total,omitempty"`
}

func (o DerivativeIndexVoSearchResultDataValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DerivativeIndexVoSearchResultDataValue struct{}"
	}

	return strings.Join([]string{"DerivativeIndexVoSearchResultDataValue", string(data)}, " ")
}
