package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAvalibleRdsResponse Response Object
type ShowAvalibleRdsResponse struct {

	// 可用目标DN。
	TargetDataNodes *[]TargetDn4Restore `json:"target_data_nodes,omitempty"`
	HttpStatusCode  int                 `json:"-"`
}

func (o ShowAvalibleRdsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAvalibleRdsResponse struct{}"
	}

	return strings.Join([]string{"ShowAvalibleRdsResponse", string(data)}, " ")
}
