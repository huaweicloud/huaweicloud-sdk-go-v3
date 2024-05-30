package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSubjectNewResponse Response Object
type UpdateSubjectNewResponse struct {
	Data           *CreateSubjectResultData `json:"data,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o UpdateSubjectNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubjectNewResponse struct{}"
	}

	return strings.Join([]string{"UpdateSubjectNewResponse", string(data)}, " ")
}
