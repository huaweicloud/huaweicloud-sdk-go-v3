package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExtendRelationId 资源归属企业项目ID
type ExtendRelationId struct {
}

func (o ExtendRelationId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtendRelationId struct{}"
	}

	return strings.Join([]string{"ExtendRelationId", string(data)}, " ")
}
