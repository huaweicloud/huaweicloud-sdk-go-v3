package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateProjectStatusRequestBody
type UpdateProjectStatusRequestBody struct {
	Project *UpdateProjectOption `json:"project"`
}

func (o UpdateProjectStatusRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProjectStatusRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateProjectStatusRequestBody", string(data)}, " ")
}
