package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSubjectResponse Response Object
type CreateSubjectResponse struct {
	Data           *CreateSubjectResultData `json:"data,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o CreateSubjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubjectResponse struct{}"
	}

	return strings.Join([]string{"CreateSubjectResponse", string(data)}, " ")
}
