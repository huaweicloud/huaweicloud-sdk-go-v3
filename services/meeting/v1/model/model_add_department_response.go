package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddDepartmentResponse Response Object
type AddDepartmentResponse struct {

	// 返回结果。
	Value          *string `json:"value,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddDepartmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDepartmentResponse struct{}"
	}

	return strings.Join([]string{"AddDepartmentResponse", string(data)}, " ")
}
