package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowShipperParamRequest Request Object
type ShowShipperParamRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 参数类型
	ParamType string `json:"param_type"`

	// 根据param的传参决定，传类型的一级菜单id，如：(lts_group_id)是一串id
	Param *string `json:"param,omitempty"`
}

func (o ShowShipperParamRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowShipperParamRequest struct{}"
	}

	return strings.Join([]string{"ShowShipperParamRequest", string(data)}, " ")
}
