package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListNodeLabelRequest struct {

	// 节点id
	ServerId string `json:"server_id"`
}

func (o ListNodeLabelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNodeLabelRequest struct{}"
	}

	return strings.Join([]string{"ListNodeLabelRequest", string(data)}, " ")
}
