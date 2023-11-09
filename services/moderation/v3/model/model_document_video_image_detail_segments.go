package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DocumentVideoImageDetailSegments struct {

	// 命中的风险片段
	Segment *string `json:"segment,omitempty"`
}

func (o DocumentVideoImageDetailSegments) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DocumentVideoImageDetailSegments struct{}"
	}

	return strings.Join([]string{"DocumentVideoImageDetailSegments", string(data)}, " ")
}
