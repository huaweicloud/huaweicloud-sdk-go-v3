package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FindingId 要检索的结果的ID。
type FindingId struct {
}

func (o FindingId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FindingId struct{}"
	}

	return strings.Join([]string{"FindingId", string(data)}, " ")
}
