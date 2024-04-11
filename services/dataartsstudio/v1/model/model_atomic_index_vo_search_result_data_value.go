package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AtomicIndexVoSearchResultDataValue 返回的数据信息。
type AtomicIndexVoSearchResultDataValue struct {

	// AtomicIndexVO数组。
	Records *[]AtomicIndexVo `json:"records,omitempty"`

	// 总数。
	Total *int32 `json:"total,omitempty"`
}

func (o AtomicIndexVoSearchResultDataValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AtomicIndexVoSearchResultDataValue struct{}"
	}

	return strings.Join([]string{"AtomicIndexVoSearchResultDataValue", string(data)}, " ")
}
