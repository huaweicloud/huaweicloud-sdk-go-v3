package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsapUuid UUID
type IsapUuid struct {
}

func (o IsapUuid) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsapUuid struct{}"
	}

	return strings.Join([]string{"IsapUuid", string(data)}, " ")
}
