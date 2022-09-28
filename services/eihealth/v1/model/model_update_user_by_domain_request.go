package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateUserByDomainRequest struct {

	// 用户id
	UserId string `json:"user_id"`

	Body *UpdateUserByDomainReq `json:"body,omitempty"`
}

func (o UpdateUserByDomainRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserByDomainRequest struct{}"
	}

	return strings.Join([]string{"UpdateUserByDomainRequest", string(data)}, " ")
}
