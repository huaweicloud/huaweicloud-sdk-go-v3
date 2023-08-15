package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTemplateGroupRequest Request Object
type DeleteTemplateGroupRequest struct {

	// 模板组id
	GroupId string `json:"group_id"`
}

func (o DeleteTemplateGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTemplateGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteTemplateGroupRequest", string(data)}, " ")
}
