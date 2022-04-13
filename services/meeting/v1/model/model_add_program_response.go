package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AddProgramResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddProgramResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddProgramResponse struct{}"
	}

	return strings.Join([]string{"AddProgramResponse", string(data)}, " ")
}
