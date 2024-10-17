package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachedEnterpriseRouterId 被挂载的企业路由器ID。
type AttachedEnterpriseRouterId struct {

	// 被挂载的企业路由器ID。
	AttachedErId string `json:"attached_er_id"`
}

func (o AttachedEnterpriseRouterId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachedEnterpriseRouterId struct{}"
	}

	return strings.Join([]string{"AttachedEnterpriseRouterId", string(data)}, " ")
}
