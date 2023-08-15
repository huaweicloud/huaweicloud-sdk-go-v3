package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNodeCertRequest Request Object
type UpdateNodeCertRequest struct {

	// 设备ID，从专业版HiLens控制台设备管理[查询设备列表](ListNodeUsingGeT.xml)获取
	NodeId string `json:"node_id"`
}

func (o UpdateNodeCertRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNodeCertRequest struct{}"
	}

	return strings.Join([]string{"UpdateNodeCertRequest", string(data)}, " ")
}
