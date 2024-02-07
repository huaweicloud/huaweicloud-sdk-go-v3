package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GenerateSurfacePointsRequest Request Object
type GenerateSurfacePointsRequest struct {

	// 平台项目ID。
	EihealthProjectId string `json:"eihealth_project_id"`

	Body *RunSurfacePointsReq `json:"body,omitempty"`
}

func (o GenerateSurfacePointsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GenerateSurfacePointsRequest struct{}"
	}

	return strings.Join([]string{"GenerateSurfacePointsRequest", string(data)}, " ")
}
