package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RelationId 关联编号
type RelationId struct {
}

func (o RelationId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RelationId struct{}"
	}

	return strings.Join([]string{"RelationId", string(data)}, " ")
}
