package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DocumentQueryResponseResultSegments struct {

	// 命中的敏感词
	Segment *string `json:"segment,omitempty"`
}

func (o DocumentQueryResponseResultSegments) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DocumentQueryResponseResultSegments struct{}"
	}

	return strings.Join([]string{"DocumentQueryResponseResultSegments", string(data)}, " ")
}
