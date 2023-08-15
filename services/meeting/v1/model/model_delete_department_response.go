package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDepartmentResponse Response Object
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
