package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TmsKey TMS标签键规范。
type TmsKey struct {
}

func (o TmsKey) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TmsKey struct{}"
	}

	return strings.Join([]string{"TmsKey", string(data)}, " ")
}
