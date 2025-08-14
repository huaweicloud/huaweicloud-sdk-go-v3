package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssignmentConfigDto struct {

	// 应用程序是否需要分配
	AssignmentRequired bool `json:"assignment_required"`
}

func (o AssignmentConfigDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssignmentConfigDto struct{}"
	}

	return strings.Join([]string{"AssignmentConfigDto", string(data)}, " ")
}
