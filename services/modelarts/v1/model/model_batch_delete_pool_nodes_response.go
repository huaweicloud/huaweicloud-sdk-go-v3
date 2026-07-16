package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeletePoolNodesResponse Response Object
type BatchDeletePoolNodesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeletePoolNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeletePoolNodesResponse struct{}"
	}

	return strings.Join([]string{"BatchDeletePoolNodesResponse", string(data)}, " ")
}
