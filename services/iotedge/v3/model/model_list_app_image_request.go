package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppImageRequest Request Object
type ListAppImageRequest struct {

	// 应用ID
	AppId string `json:"app_id"`

	// 应用版本
	Version string `json:"version"`

	// 每页记录数，默认值为10，取值区间为1-1000。
	Limit *int32 `json:"limit,omitempty"`

	// 查询的起始位置，取值范围为非负整数，默认为0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListAppImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppImageRequest struct{}"
	}

	return strings.Join([]string{"ListAppImageRequest", string(data)}, " ")
}
