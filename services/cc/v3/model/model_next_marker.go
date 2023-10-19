package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NextMarker 下一页的marker，为空时表示最后一页。
type NextMarker struct {

	// 向后分页标识符。
	NextMarker *string `json:"next_marker,omitempty"`
}

func (o NextMarker) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NextMarker struct{}"
	}

	return strings.Join([]string{"NextMarker", string(data)}, " ")
}
