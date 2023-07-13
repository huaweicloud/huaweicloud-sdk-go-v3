package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePubInfoResponseModelData 响应实体类。
type CreatePubInfoResponseModelData struct {

	// 服务号申请记录ID。
	LogId string `json:"log_id"`

	// 服务号ID。
	PubId *string `json:"pub_id,omitempty"`

	// 菜单ID。
	MenuId *string `json:"menu_id,omitempty"`

	// 主页ID。
	PortalId *string `json:"portal_id,omitempty"`

	// 服务号名称。
	PubName *string `json:"pub_name,omitempty"`
}

func (o CreatePubInfoResponseModelData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePubInfoResponseModelData struct{}"
	}

	return strings.Join([]string{"CreatePubInfoResponseModelData", string(data)}, " ")
}
