package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscribeDataRequest Request Object
type SubscribeDataRequest struct {

	// 资产订阅目标项目
	EihealthProjectId string `json:"eihealth_project_id"`

	Body *SubscribeDataReq `json:"body,omitempty"`
}

func (o SubscribeDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscribeDataRequest struct{}"
	}

	return strings.Join([]string{"SubscribeDataRequest", string(data)}, " ")
}
