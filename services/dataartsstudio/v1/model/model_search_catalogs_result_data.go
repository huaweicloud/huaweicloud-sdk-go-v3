package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchCatalogsResultData data，统一的返回结果的最外层数据结构。
type SearchCatalogsResultData struct {
	Value *SearchCatalogsResultDataValue `json:"value,omitempty"`
}

func (o SearchCatalogsResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchCatalogsResultData struct{}"
	}

	return strings.Join([]string{"SearchCatalogsResultData", string(data)}, " ")
}
