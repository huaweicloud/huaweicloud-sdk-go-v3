package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSessionTypeResponse Response Object
type ListSessionTypeResponse struct {

	// 会话列表
	SessionTypes   *[]SessionTypeEntity `json:"session_types,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListSessionTypeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSessionTypeResponse struct{}"
	}

	return strings.Join([]string{"ListSessionTypeResponse", string(data)}, " ")
}
