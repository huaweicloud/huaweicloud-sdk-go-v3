package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAutoKillProcessStatusResponse Response Object
type UpdateAutoKillProcessStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateAutoKillProcessStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAutoKillProcessStatusResponse struct{}"
	}

	return strings.Join([]string{"UpdateAutoKillProcessStatusResponse", string(data)}, " ")
}
