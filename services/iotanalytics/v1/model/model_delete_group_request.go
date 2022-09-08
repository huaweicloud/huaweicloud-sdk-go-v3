package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteGroupRequest struct {

	// 存储组 ID
	GroupId string `json:"group_id"`
}

func (o DeleteGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteGroupRequest", string(data)}, " ")
}
