package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsapTableId 表ID
type IsapTableId struct {
}

func (o IsapTableId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsapTableId struct{}"
	}

	return strings.Join([]string{"IsapTableId", string(data)}, " ")
}
