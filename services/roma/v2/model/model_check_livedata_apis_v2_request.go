package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CheckLivedataApisV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *LdApiCheckInfo `json:"body,omitempty" xml:"body"`
}

func (o CheckLivedataApisV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckLivedataApisV2Request struct{}"
	}

	return strings.Join([]string{"CheckLivedataApisV2Request", string(data)}, " ")
}
