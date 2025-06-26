package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteResourceTagsRequestBody struct {

	// 批量操作标签
	Tags []ResourceTag `json:"tags"`
}

func (o DeleteResourceTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResourceTagsRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteResourceTagsRequestBody", string(data)}, " ")
}
