package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssociateAppsRequestBody struct {

	// 所需要绑定的应用ID
	AppId string `json:"app_id"`

	// 所需要绑定的集群ID
	ClusterId string `json:"cluster_id"`
}

func (o AssociateAppsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateAppsRequestBody struct{}"
	}

	return strings.Join([]string{"AssociateAppsRequestBody", string(data)}, " ")
}
