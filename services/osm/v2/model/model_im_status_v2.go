package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImStatusV2 struct {

	// 状态
	Status *int32 `json:"status,omitempty" xml:"status"`

	// 工单id
	IncidentId *string `json:"incident_id,omitempty" xml:"incident_id"`
}

func (o ImStatusV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImStatusV2 struct{}"
	}

	return strings.Join([]string{"ImStatusV2", string(data)}, " ")
}
