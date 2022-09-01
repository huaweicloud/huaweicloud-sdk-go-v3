package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowTaskSetResponse struct {

	// code
	Code *string `json:"code,omitempty" xml:"code"`

	// extend
	Extend *[]string `json:"extend,omitempty" xml:"extend"`

	// message
	Message *string `json:"message,omitempty" xml:"message"`

	// 工程集详细信息
	Tasks          *[]Task `json:"tasks,omitempty" xml:"tasks"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTaskSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskSetResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskSetResponse", string(data)}, " ")
}
