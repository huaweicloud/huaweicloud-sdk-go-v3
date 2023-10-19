package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ParseUserBehaviorRequest Request Object
type ParseUserBehaviorRequest struct {

	// 实例id
	Instance string `json:"instance"`

	Body *BehaviorRestBody `json:"body,omitempty"`
}

func (o ParseUserBehaviorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParseUserBehaviorRequest struct{}"
	}

	return strings.Join([]string{"ParseUserBehaviorRequest", string(data)}, " ")
}
