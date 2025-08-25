package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DisAssociateAppsRequestBody struct {

	// 所需要绑定的集群ID
	ClusterId string `json:"cluster_id"`

	// 所需要解绑的应用ID列表
	AppIds []string `json:"app_ids"`
}

func (o DisAssociateAppsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisAssociateAppsRequestBody struct{}"
	}

	return strings.Join([]string{"DisAssociateAppsRequestBody", string(data)}, " ")
}
