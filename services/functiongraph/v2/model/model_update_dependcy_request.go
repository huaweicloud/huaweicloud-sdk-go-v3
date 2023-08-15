package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDependcyRequest Request Object
type UpdateDependcyRequest struct {

	// 依赖包的ID。
	DependId string `json:"depend_id"`

	Body *UpdateDependencyRequestBody `json:"body,omitempty"`
}

func (o UpdateDependcyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDependcyRequest struct{}"
	}

	return strings.Join([]string{"UpdateDependcyRequest", string(data)}, " ")
}
