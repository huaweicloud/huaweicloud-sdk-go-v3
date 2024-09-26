package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GcbProjectId struct {

	// 功能说明：实例所在region对应的projectId。
	ProjectId string `json:"project_id"`
}

func (o GcbProjectId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GcbProjectId struct{}"
	}

	return strings.Join([]string{"GcbProjectId", string(data)}, " ")
}
