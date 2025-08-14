package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAssessTaskResponse Response Object
type CreateAssessTaskResponse struct {
	Data           *BaseIdResponseData `json:"data,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o CreateAssessTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAssessTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateAssessTaskResponse", string(data)}, " ")
}
