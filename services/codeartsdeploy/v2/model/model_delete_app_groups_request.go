package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAppGroupsRequest Request Object
type DeleteAppGroupsRequest struct {

	// 项目Id
	ProjectId string `json:"project_id"`

	// 分组Id
	GroupId string `json:"group_id"`
}

func (o DeleteAppGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAppGroupsRequest struct{}"
	}

	return strings.Join([]string{"DeleteAppGroupsRequest", string(data)}, " ")
}
