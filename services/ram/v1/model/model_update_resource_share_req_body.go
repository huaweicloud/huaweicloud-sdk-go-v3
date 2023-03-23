package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// The request body of the UpdateResourceShare operation.
type UpdateResourceShareReqBody struct {

	// 资源共享实例的名称。
	Name string `json:"name"`

	// 资源共享实例的描述。
	Description *string `json:"description,omitempty"`
}

func (o UpdateResourceShareReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResourceShareReqBody struct{}"
	}

	return strings.Join([]string{"UpdateResourceShareReqBody", string(data)}, " ")
}
