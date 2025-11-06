package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIntelligentKillSessionRequest Request Object
type CreateIntelligentKillSessionRequest struct {

	// 实例id
	InstanceId string `json:"instance_id"`

	Body *IntelligentKillSessionReq `json:"body,omitempty"`
}

func (o CreateIntelligentKillSessionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIntelligentKillSessionRequest struct{}"
	}

	return strings.Join([]string{"CreateIntelligentKillSessionRequest", string(data)}, " ")
}
