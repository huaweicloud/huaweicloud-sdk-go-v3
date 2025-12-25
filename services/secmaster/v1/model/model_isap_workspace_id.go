package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsapWorkspaceId 工作空间ID
type IsapWorkspaceId struct {
}

func (o IsapWorkspaceId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsapWorkspaceId struct{}"
	}

	return strings.Join([]string{"IsapWorkspaceId", string(data)}, " ")
}
