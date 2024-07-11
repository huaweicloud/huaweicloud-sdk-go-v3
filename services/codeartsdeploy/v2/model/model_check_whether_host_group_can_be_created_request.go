package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckWhetherHostGroupCanBeCreatedRequest Request Object
type CheckWhetherHostGroupCanBeCreatedRequest struct {

	// 项目id
	ProjectId string `json:"project_id"`
}

func (o CheckWhetherHostGroupCanBeCreatedRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckWhetherHostGroupCanBeCreatedRequest struct{}"
	}

	return strings.Join([]string{"CheckWhetherHostGroupCanBeCreatedRequest", string(data)}, " ")
}
