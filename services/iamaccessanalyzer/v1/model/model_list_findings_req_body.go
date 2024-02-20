package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListFindingsReqBody struct {
	Filters *[]FindingFilter `json:"filters,omitempty"`

	// 单页最大结果数。
	Limit *int32 `json:"limit,omitempty"`

	// 页面标记。
	Marker *string `json:"marker,omitempty"`
}

func (o ListFindingsReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFindingsReqBody struct{}"
	}

	return strings.Join([]string{"ListFindingsReqBody", string(data)}, " ")
}
