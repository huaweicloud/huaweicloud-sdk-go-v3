package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteProgramsResponse Response Object
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
