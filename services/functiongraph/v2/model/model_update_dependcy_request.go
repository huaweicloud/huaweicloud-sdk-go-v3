package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateDependcyRequest struct {

	// 依赖包的ID。
	DependId string `json:"depend_id" xml:"depend_id"`

	Body *UpdateDependencyRequestBody `json:"body,omitempty" xml:"body"`
}

func (o UpdateDependcyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDependcyRequest struct{}"
	}

	return strings.Join([]string{"UpdateDependcyRequest", string(data)}, " ")
}
