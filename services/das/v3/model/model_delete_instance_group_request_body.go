package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteInstanceGroupRequestBody struct {

	// 实例组ID
	GroupId int32 `json:"group_id"`
}

func (o DeleteInstanceGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceGroupRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteInstanceGroupRequestBody", string(data)}, " ")
}
