package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckPermissionInput CheckPermissionInput body
type CheckPermissionInput struct {

	// 主体信息
	AccessRequest []AccessRequest `json:"access_request"`
}

func (o CheckPermissionInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckPermissionInput struct{}"
	}

	return strings.Join([]string{"CheckPermissionInput", string(data)}, " ")
}
