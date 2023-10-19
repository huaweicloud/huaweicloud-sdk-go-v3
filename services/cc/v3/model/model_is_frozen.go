package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsFrozen 是否冻结
type IsFrozen struct {

	// 是否冻结
	IsFrozen bool `json:"is_frozen"`
}

func (o IsFrozen) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsFrozen struct{}"
	}

	return strings.Join([]string{"IsFrozen", string(data)}, " ")
}
