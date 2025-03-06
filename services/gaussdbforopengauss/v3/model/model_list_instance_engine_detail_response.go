package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceEngineDetailResponse Response Object
type ListInstanceEngineDetailResponse struct {

	// 引擎版本和相应的实例详情。
	EngineInstanceDetails *[]InstanceEngineDetail `json:"engine_instance_details,omitempty"`

	// 引擎版本数量。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstanceEngineDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceEngineDetailResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceEngineDetailResponse", string(data)}, " ")
}
