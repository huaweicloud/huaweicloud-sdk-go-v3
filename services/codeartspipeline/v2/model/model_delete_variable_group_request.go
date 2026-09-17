package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVariableGroupRequest Request Object
type DeleteVariableGroupRequest struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	// 参数组ID
	Id string `json:"id"`
}

func (o DeleteVariableGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVariableGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteVariableGroupRequest", string(data)}, " ")
}
