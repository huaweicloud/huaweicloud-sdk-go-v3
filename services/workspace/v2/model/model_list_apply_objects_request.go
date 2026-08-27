package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApplyObjectsRequest Request Object
type ListApplyObjectsRequest struct {

	// 策略ID（精确查询）
	StrategyId string `json:"strategy_id"`

	// 应用对象名称（支持模糊查询）
	ObjectName *string `json:"object_name,omitempty"`

	// 应用对象类型，包括DESKTOP、ALL_DESKTOPS、DESKTOP_POOL、DESKTOP_TAG、ALL_USERS、USER、USER_GROUP
	ObjectType *string `json:"object_type,omitempty"`

	// 偏移量，默认0
	Offset *int32 `json:"offset,omitempty"`

	// 每页数量，默认10，最大100
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListApplyObjectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplyObjectsRequest struct{}"
	}

	return strings.Join([]string{"ListApplyObjectsRequest", string(data)}, " ")
}
