package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListVersionRequest struct {
	// 待查询版本号。当前仅支持v2。

	ApiVersion string `json:"api_version"`
}

func (o ListVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVersionRequest struct{}"
	}

	return strings.Join([]string{"ListVersionRequest", string(data)}, " ")
}
