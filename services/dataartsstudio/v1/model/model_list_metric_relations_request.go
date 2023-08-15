package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMetricRelationsRequest Request Object
type ListMetricRelationsRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 实体id
	Id string `json:"id"`

	// 指标类型
	BizType string `json:"biz_type"`
}

func (o ListMetricRelationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetricRelationsRequest struct{}"
	}

	return strings.Join([]string{"ListMetricRelationsRequest", string(data)}, " ")
}
