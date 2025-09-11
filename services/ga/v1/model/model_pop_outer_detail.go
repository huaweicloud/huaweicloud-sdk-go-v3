package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PopOuterDetail 接入点详情。
type PopOuterDetail struct {

	// 接入点ID。
	Id *string `json:"id,omitempty"`
}

func (o PopOuterDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PopOuterDetail struct{}"
	}

	return strings.Join([]string{"PopOuterDetail", string(data)}, " ")
}
