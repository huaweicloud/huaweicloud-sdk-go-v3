package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteResourceShareRequest Request Object
type DeleteResourceShareRequest struct {

	// 资源共享实例的ID。
	ResourceShareId string `json:"resource_share_id"`
}

func (o DeleteResourceShareRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResourceShareRequest struct{}"
	}

	return strings.Join([]string{"DeleteResourceShareRequest", string(data)}, " ")
}
