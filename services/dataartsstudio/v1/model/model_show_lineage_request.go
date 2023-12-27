package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLineageRequest Request Object
type ShowLineageRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 资产guid
	Guid string `json:"guid"`

	// 查询方向，取值范围：BOTH、IN、OUT。默认BOTH
	Direction *string `json:"direction,omitempty"`

	// 血缘链路长度，默认值5
	Depth *int32 `json:"depth,omitempty"`
}

func (o ShowLineageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLineageRequest struct{}"
	}

	return strings.Join([]string{"ShowLineageRequest", string(data)}, " ")
}
