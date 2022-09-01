package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteEdgeApplicationVersionRequest struct {

	// 应用版本,应用内版本唯一。
	EdgeAppId string `json:"edge_app_id" xml:"edge_app_id"`

	// 应用版本ID，应用版本唯一。
	Version string `json:"version" xml:"version"`
}

func (o DeleteEdgeApplicationVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEdgeApplicationVersionRequest struct{}"
	}

	return strings.Join([]string{"DeleteEdgeApplicationVersionRequest", string(data)}, " ")
}
