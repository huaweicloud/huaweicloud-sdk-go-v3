package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDepartmentResponse Response Object
type UpdateDepartmentResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateDepartmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDepartmentResponse struct{}"
	}

	return strings.Join([]string{"UpdateDepartmentResponse", string(data)}, " ")
}
