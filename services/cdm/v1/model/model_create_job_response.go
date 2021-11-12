package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateJobResponse struct {
	// 作业名称。

	Name *string `json:"name,omitempty"`
	// 校验结果： - 如果修改失败，返回失败原因。 - 如果修改成功，返回空列表。

	ValidationResult *[]JobValidationResult `json:"validation-result,omitempty"`
	HttpStatusCode   int                    `json:"-"`
}

func (o CreateJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateJobResponse struct{}"
	}

	return strings.Join([]string{"CreateJobResponse", string(data)}, " ")
}
