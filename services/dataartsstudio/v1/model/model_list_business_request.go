package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBusinessRequest Request Object
type ListBusinessRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 查询条数，即查询Y条数据。默认值50，取值范围[1,100]
	Limit *int32 `json:"limit,omitempty"`

	// 查询起始坐标，即跳过X条数据，仅支持0或limit的整数倍，不满足则向下取整。默认值0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListBusinessRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBusinessRequest struct{}"
	}

	return strings.Join([]string{"ListBusinessRequest", string(data)}, " ")
}
