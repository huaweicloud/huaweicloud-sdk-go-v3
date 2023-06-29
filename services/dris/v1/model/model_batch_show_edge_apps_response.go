package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchShowEdgeAppsResponse Response Object
type BatchShowEdgeAppsResponse struct {

	// **参数说明**：总记录数。
	Count *int64 `json:"count,omitempty"`

	// **参数说明**：列举每条记录。
	EdgeApps       *[]QueryApplicationBriefResponseDto `json:"edge_apps,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o BatchShowEdgeAppsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchShowEdgeAppsResponse struct{}"
	}

	return strings.Join([]string{"BatchShowEdgeAppsResponse", string(data)}, " ")
}
