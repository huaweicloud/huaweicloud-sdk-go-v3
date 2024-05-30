package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSubjectNewResponse Response Object
type CreateSubjectNewResponse struct {
	Data           *CreateSubjectResultData `json:"data,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o CreateSubjectNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubjectNewResponse struct{}"
	}

	return strings.Join([]string{"CreateSubjectNewResponse", string(data)}, " ")
}
