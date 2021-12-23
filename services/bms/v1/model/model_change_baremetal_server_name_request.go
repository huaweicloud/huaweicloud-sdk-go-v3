package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ChangeBaremetalServerNameRequest struct {
	// 裸金属服务器ID

	ServerId string `json:"server_id"`

	Body *ChangeBaremetalNameBody `json:"body,omitempty"`
}

func (o ChangeBaremetalServerNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeBaremetalServerNameRequest struct{}"
	}

	return strings.Join([]string{"ChangeBaremetalServerNameRequest", string(data)}, " ")
}
