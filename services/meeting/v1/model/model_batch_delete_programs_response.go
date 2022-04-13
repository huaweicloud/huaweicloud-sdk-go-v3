package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchDeleteProgramsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteProgramsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteProgramsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteProgramsResponse", string(data)}, " ")
}
