package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CancelAutoRenewalResourcesRequest struct {
	// 资源ID。 您可以调用“查询客户包年/包月资源列表”接口获取资源ID。 设置主资源时会将从资源一起设置，主从关系为： 云主机为主资源，对应的从资源为云硬盘共享带宽的情况下，带宽为主资源，对应的从资源为弹性IP（可能包含多个IP）独享带宽的情况下，弹性IP为主资源，对应的从资源为带宽

	ResourceId string `json:"resource_id"`
}

func (o CancelAutoRenewalResourcesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CancelAutoRenewalResourcesRequest struct{}"
	}

	return strings.Join([]string{"CancelAutoRenewalResourcesRequest", string(data)}, " ")
}
