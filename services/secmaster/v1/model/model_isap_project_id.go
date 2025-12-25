package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsapProjectId 项目ID
type IsapProjectId struct {
}

func (o IsapProjectId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsapProjectId struct{}"
	}

	return strings.Join([]string{"IsapProjectId", string(data)}, " ")
}
