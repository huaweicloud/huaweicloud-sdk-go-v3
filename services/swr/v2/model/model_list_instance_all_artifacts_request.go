package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceAllArtifactsRequest Request Object
type ListInstanceAllArtifactsRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 分页查询时的查询标记，使用上一次接口调用返回的next_marker值。默认值为1。**注意：marker和limit参数需要配套使用。**
	Marker *int32 `json:"marker,omitempty"`

	// 返回条数，默认为10，最大值为100。**注意：marker和limit参数需要配套使用。**
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListInstanceAllArtifactsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceAllArtifactsRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceAllArtifactsRequest", string(data)}, " ")
}
