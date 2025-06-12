package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TotalResponseBodyResult 结果
type TotalResponseBodyResult struct {

	// 总数
	Total *int32 `json:"total,omitempty"`
}

func (o TotalResponseBodyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TotalResponseBodyResult struct{}"
	}

	return strings.Join([]string{"TotalResponseBodyResult", string(data)}, " ")
}
