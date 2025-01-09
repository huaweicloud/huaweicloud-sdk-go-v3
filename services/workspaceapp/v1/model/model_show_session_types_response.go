package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSessionTypesResponse Response Object
type ShowSessionTypesResponse struct {

	// 会话列表。
	SessionTypes   *[]SessionTypeEntity `json:"session_types,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowSessionTypesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSessionTypesResponse struct{}"
	}

	return strings.Join([]string{"ShowSessionTypesResponse", string(data)}, " ")
}
