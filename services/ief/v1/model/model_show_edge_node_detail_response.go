package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowEdgeNodeDetailResponse struct {
	Node           *EdgeNodeResp `json:"node,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowEdgeNodeDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEdgeNodeDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowEdgeNodeDetailResponse", string(data)}, " ")
}
