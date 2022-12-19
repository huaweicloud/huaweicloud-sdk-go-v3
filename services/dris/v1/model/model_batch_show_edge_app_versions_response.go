package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchShowEdgeAppVersionsResponse struct {

	// **参数说明**：总记录数。
	Count *int64 `json:"count,omitempty"`

	// **参数说明**：列举每条记录。
	Versions       *[]QueryEdgeAppVersionBriefResponseDto `json:"versions,omitempty"`
	HttpStatusCode int                                    `json:"-"`
}

func (o BatchShowEdgeAppVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchShowEdgeAppVersionsResponse struct{}"
	}

	return strings.Join([]string{"BatchShowEdgeAppVersionsResponse", string(data)}, " ")
}
