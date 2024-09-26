package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGroupRequest Request Object
type ShowGroupRequest struct {

	// 项目id
	ProjectId string `json:"project_id"`

	// 代码组id
	GroupId int32 `json:"group_id"`
}

func (o ShowGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowGroupRequest", string(data)}, " ")
}
