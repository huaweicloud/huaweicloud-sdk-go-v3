package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodeRequest struct {
	Node *NodeReqDetail `json:"node"`
}

func (o NodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeRequest struct{}"
	}

	return strings.Join([]string{"NodeRequest", string(data)}, " ")
}
