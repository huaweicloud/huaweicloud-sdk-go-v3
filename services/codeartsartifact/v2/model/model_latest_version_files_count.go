package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LatestVersionFilesCount struct {

	// **参数解释**: 数目。 **取值范围**: 不涉及。
	Count *int64 `json:"count,omitempty"`
}

func (o LatestVersionFilesCount) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LatestVersionFilesCount struct{}"
	}

	return strings.Join([]string{"LatestVersionFilesCount", string(data)}, " ")
}
