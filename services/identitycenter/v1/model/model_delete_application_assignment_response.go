package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteApplicationAssignmentResponse Response Object
type DeleteApplicationAssignmentResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteApplicationAssignmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteApplicationAssignmentResponse struct{}"
	}

	return strings.Join([]string{"DeleteApplicationAssignmentResponse", string(data)}, " ")
}
