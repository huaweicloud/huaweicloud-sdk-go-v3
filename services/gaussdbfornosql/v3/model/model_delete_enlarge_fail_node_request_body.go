package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteEnlargeFailNodeRequestBody struct {

	// 节点ID
	NodeId string `json:"node_id"`
}

func (o DeleteEnlargeFailNodeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEnlargeFailNodeRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteEnlargeFailNodeRequestBody", string(data)}, " ")
}
