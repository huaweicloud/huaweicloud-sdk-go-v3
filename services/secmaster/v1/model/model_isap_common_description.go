package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsapCommonDescription 描述信息
type IsapCommonDescription struct {
}

func (o IsapCommonDescription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsapCommonDescription struct{}"
	}

	return strings.Join([]string{"IsapCommonDescription", string(data)}, " ")
}
