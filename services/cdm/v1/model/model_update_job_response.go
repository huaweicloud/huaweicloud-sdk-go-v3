package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateJobResponse struct {
	// 校验结果：如果修改失败，返回失败原因。如果修改成功，返回空列表

	ValidationResult *[]JobValidationResult `json:"validation-result,omitempty"`
	HttpStatusCode   int                    `json:"-"`
}

func (o UpdateJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateJobResponse struct{}"
	}

	return strings.Join([]string{"UpdateJobResponse", string(data)}, " ")
}
