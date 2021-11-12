package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateDependencyRequest struct {
	// 依赖包的ID。

	DependId string `json:"depend_id"`

	Body *UpdateDependencyRequestBody `json:"body,omitempty"`
}

func (o UpdateDependencyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDependencyRequest struct{}"
	}

	return strings.Join([]string{"UpdateDependencyRequest", string(data)}, " ")
}
