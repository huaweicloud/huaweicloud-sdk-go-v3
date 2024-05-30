package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeSubjectsResponse Response Object
type ChangeSubjectsResponse struct {
	Data           *ChangeSubjectsResultData `json:"data,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ChangeSubjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeSubjectsResponse struct{}"
	}

	return strings.Join([]string{"ChangeSubjectsResponse", string(data)}, " ")
}
