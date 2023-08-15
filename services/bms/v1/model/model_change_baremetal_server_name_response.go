package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeBaremetalServerNameResponse Response Object
type ChangeBaremetalServerNameResponse struct {
	Server         *ChangeBaremetalNameResponsesServers `json:"server,omitempty"`
	HttpStatusCode int                                  `json:"-"`
}

func (o ChangeBaremetalServerNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeBaremetalServerNameResponse struct{}"
	}

	return strings.Join([]string{"ChangeBaremetalServerNameResponse", string(data)}, " ")
}
