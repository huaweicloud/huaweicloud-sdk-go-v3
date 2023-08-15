package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteAppGroupResponse Response Object
type BatchDeleteAppGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteAppGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAppGroupResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteAppGroupResponse", string(data)}, " ")
}
