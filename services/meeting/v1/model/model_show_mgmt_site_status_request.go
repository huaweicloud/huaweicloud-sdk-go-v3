package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMgmtSiteStatusRequest Request Object
type ShowMgmtSiteStatusRequest struct {

	// 模版类型。
	XSchemaType *string `json:"X-Schema-Type,omitempty"`
}

func (o ShowMgmtSiteStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMgmtSiteStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowMgmtSiteStatusRequest", string(data)}, " ")
}
