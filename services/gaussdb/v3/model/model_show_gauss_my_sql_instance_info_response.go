package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowGaussMySqlInstanceInfoResponse struct {
	Instance       *MysqlInstanceInfoDetail `json:"instance,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ShowGaussMySqlInstanceInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGaussMySqlInstanceInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowGaussMySqlInstanceInfoResponse", string(data)}, " ")
}
