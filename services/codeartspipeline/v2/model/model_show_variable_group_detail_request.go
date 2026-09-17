package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVariableGroupDetailRequest Request Object
type ShowVariableGroupDetailRequest struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	// 参数组ID
	Id string `json:"id"`
}

func (o ShowVariableGroupDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVariableGroupDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowVariableGroupDetailRequest", string(data)}, " ")
}
