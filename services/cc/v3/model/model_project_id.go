package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProjectId 实例所属项目ID。
type ProjectId struct {

	// 实例所属项目ID。
	ProjectId string `json:"project_id"`
}

func (o ProjectId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectId struct{}"
	}

	return strings.Join([]string{"ProjectId", string(data)}, " ")
}
