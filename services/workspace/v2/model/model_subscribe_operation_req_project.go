package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscribeOperationReqProject 项目信息
type SubscribeOperationReqProject struct {

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`
}

func (o SubscribeOperationReqProject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscribeOperationReqProject struct{}"
	}

	return strings.Join([]string{"SubscribeOperationReqProject", string(data)}, " ")
}
