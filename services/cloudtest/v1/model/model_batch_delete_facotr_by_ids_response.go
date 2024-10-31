package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteFacotrByIdsResponse Response Object
type BatchDeleteFacotrByIdsResponse struct {
	Code *string `json:"code,omitempty"`

	Data *interface{} `json:"data,omitempty"`

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchDeleteFacotrByIdsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteFacotrByIdsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteFacotrByIdsResponse", string(data)}, " ")
}
