package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteEdgeAppRequest struct {
	// 应用ID，应用唯一。

	EdgeAppId string `json:"edge_app_id"`
}

func (o DeleteEdgeAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEdgeAppRequest struct{}"
	}

	return strings.Join([]string{"DeleteEdgeAppRequest", string(data)}, " ")
}
