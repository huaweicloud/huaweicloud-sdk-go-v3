package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateConstraintResponse Response Object
type BatchCreateConstraintResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateConstraintResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateConstraintResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateConstraintResponse", string(data)}, " ")
}
