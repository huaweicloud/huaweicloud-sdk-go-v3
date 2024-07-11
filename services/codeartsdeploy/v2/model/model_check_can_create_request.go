package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckCanCreateRequest Request Object
type CheckCanCreateRequest struct {

	// 项目id
	ProjectId string `json:"project_id"`
}

func (o CheckCanCreateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckCanCreateRequest struct{}"
	}

	return strings.Join([]string{"CheckCanCreateRequest", string(data)}, " ")
}
