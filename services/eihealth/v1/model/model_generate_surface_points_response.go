package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GenerateSurfacePointsResponse Response Object
type GenerateSurfacePointsResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o GenerateSurfacePointsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GenerateSurfacePointsResponse struct{}"
	}

	return strings.Join([]string{"GenerateSurfacePointsResponse", string(data)}, " ")
}
