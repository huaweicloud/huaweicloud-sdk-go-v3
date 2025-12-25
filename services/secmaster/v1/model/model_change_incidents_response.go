package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeIncidentsResponse Response Object
type ChangeIncidentsResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	Data           *IncidentDetail `json:"data,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ChangeIncidentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeIncidentsResponse struct{}"
	}

	return strings.Join([]string{"ChangeIncidentsResponse", string(data)}, " ")
}
