package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateProxyPrivateDnsNameResponse Response Object
type UpdateProxyPrivateDnsNameResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateProxyPrivateDnsNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProxyPrivateDnsNameResponse struct{}"
	}

	return strings.Join([]string{"UpdateProxyPrivateDnsNameResponse", string(data)}, " ")
}
