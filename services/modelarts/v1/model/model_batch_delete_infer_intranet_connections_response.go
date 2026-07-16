package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteInferIntranetConnectionsResponse Response Object
type BatchDeleteInferIntranetConnectionsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteInferIntranetConnectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteInferIntranetConnectionsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteInferIntranetConnectionsResponse", string(data)}, " ")
}
