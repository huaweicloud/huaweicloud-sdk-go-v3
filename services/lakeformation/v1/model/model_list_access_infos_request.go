package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccessInfosRequest Request Object
type ListAccessInfosRequest struct {

	// LakeFormation实例ID
	InstanceId string `json:"instance_id"`

	// 分页查询时的偏移量
	Offset int32 `json:"offset"`

	// 分页一页显示数
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAccessInfosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccessInfosRequest struct{}"
	}

	return strings.Join([]string{"ListAccessInfosRequest", string(data)}, " ")
}
