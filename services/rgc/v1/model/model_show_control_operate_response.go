package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowControlOperateResponse Response Object
type ShowControlOperateResponse struct {
	ControlOperation *ControlOperation `json:"control_operation,omitempty"`
	HttpStatusCode   int               `json:"-"`
}

func (o ShowControlOperateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowControlOperateResponse struct{}"
	}

	return strings.Join([]string{"ShowControlOperateResponse", string(data)}, " ")
}
