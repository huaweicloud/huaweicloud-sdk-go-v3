package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StandardSearchResultData 返回数据标准集合。
type StandardSearchResultData struct {
	Value *StandardSearchResultDataValue `json:"value,omitempty"`
}

func (o StandardSearchResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StandardSearchResultData struct{}"
	}

	return strings.Join([]string{"StandardSearchResultData", string(data)}, " ")
}
