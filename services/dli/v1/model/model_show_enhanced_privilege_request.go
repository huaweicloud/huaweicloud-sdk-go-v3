package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEnhancedPrivilegeRequest Request Object
type ShowEnhancedPrivilegeRequest struct {

	// 增强型跨源连接ID。
	ConnectionId string `json:"connection_id"`
}

func (o ShowEnhancedPrivilegeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEnhancedPrivilegeRequest struct{}"
	}

	return strings.Join([]string{"ShowEnhancedPrivilegeRequest", string(data)}, " ")
}
