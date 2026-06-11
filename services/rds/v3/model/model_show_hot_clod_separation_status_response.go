package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHotClodSeparationStatusResponse Response Object
type ShowHotClodSeparationStatusResponse struct {

	// 功能是否开启
	Opened         *bool `json:"opened,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowHotClodSeparationStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHotClodSeparationStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowHotClodSeparationStatusResponse", string(data)}, " ")
}
