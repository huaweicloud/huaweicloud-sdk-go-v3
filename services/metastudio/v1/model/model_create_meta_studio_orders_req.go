package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMetaStudioOrdersReq 下单请求体
type CreateMetaStudioOrdersReq struct {

	// 云服务信息列表
	CloudServices *[]PublicCloudServiceOrder `json:"cloud_services,omitempty"`
}

func (o CreateMetaStudioOrdersReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMetaStudioOrdersReq struct{}"
	}

	return strings.Join([]string{"CreateMetaStudioOrdersReq", string(data)}, " ")
}
