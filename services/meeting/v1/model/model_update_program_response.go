package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateProgramResponse Response Object
type UpdateProgramResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateProgramResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProgramResponse struct{}"
	}

	return strings.Join([]string{"UpdateProgramResponse", string(data)}, " ")
}
