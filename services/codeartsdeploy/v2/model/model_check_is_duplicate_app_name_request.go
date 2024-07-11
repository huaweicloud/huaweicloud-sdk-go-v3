package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckIsDuplicateAppNameRequest Request Object
type CheckIsDuplicateAppNameRequest struct {

	// 应用名称
	Name string `json:"name"`

	// 项目id
	ProjectId string `json:"project_id"`
}

func (o CheckIsDuplicateAppNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckIsDuplicateAppNameRequest struct{}"
	}

	return strings.Join([]string{"CheckIsDuplicateAppNameRequest", string(data)}, " ")
}
