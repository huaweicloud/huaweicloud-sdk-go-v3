package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PreviousMarker 上一页的marker，为空时表示第一页。
type PreviousMarker struct {

	// 向前分页标识符。
	PreviousMarker *string `json:"previous_marker,omitempty"`
}

func (o PreviousMarker) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PreviousMarker struct{}"
	}

	return strings.Join([]string{"PreviousMarker", string(data)}, " ")
}
