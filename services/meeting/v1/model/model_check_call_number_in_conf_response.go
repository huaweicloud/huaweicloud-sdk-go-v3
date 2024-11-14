package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckCallNumberInConfResponse Response Object
type CheckCallNumberInConfResponse struct {

	// 是否在会议中
	InConf *bool `json:"in_conf,omitempty"`

	// 会议id
	ConfId         *string `json:"conf_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CheckCallNumberInConfResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckCallNumberInConfResponse struct{}"
	}

	return strings.Join([]string{"CheckCallNumberInConfResponse", string(data)}, " ")
}
