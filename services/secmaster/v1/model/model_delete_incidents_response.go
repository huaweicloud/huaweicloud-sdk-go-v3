package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteIncidentsResponse Response Object
type DeleteIncidentsResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	Data           *DeleteIncidentResponseBodyData `json:"data,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o DeleteIncidentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIncidentsResponse struct{}"
	}

	return strings.Join([]string{"DeleteIncidentsResponse", string(data)}, " ")
}
