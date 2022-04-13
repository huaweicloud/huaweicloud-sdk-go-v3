package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteDepartmentResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDepartmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDepartmentResponse struct{}"
	}

	return strings.Join([]string{"DeleteDepartmentResponse", string(data)}, " ")
}
