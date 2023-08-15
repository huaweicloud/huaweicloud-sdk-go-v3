package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCommandResultResponse Response Object
type UpdateCommandResultResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateCommandResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCommandResultResponse struct{}"
	}

	return strings.Join([]string{"UpdateCommandResultResponse", string(data)}, " ")
}
