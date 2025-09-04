package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteProxyPrivateDnsNameResponse Response Object
type DeleteProxyPrivateDnsNameResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteProxyPrivateDnsNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProxyPrivateDnsNameResponse struct{}"
	}

	return strings.Join([]string{"DeleteProxyPrivateDnsNameResponse", string(data)}, " ")
}
