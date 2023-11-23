package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEnhancedConnectionPrivilegeRequest Request Object
type ShowEnhancedConnectionPrivilegeRequest struct {

	// 增强型跨源连接ID。
	ConnectionId string `json:"connection_id"`
}

func (o ShowEnhancedConnectionPrivilegeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEnhancedConnectionPrivilegeRequest struct{}"
	}

	return strings.Join([]string{"ShowEnhancedConnectionPrivilegeRequest", string(data)}, " ")
}
