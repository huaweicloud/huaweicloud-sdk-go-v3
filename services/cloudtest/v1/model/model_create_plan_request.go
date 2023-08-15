package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePlanRequest Request Object
type CreatePlanRequest struct {

	// 项目唯一标识，固定长度32位字符
	ProjectId string `json:"project_id"`

	Body *CreatePlanRequestBody `json:"body,omitempty"`
}

func (o CreatePlanRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePlanRequest struct{}"
	}

	return strings.Join([]string{"CreatePlanRequest", string(data)}, " ")
}
