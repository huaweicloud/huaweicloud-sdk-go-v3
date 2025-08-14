package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateApplicationAssignmentResponse Response Object
type CreateApplicationAssignmentResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateApplicationAssignmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApplicationAssignmentResponse struct{}"
	}

	return strings.Join([]string{"CreateApplicationAssignmentResponse", string(data)}, " ")
}
