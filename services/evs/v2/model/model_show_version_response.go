package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVersionResponse Response Object
type ShowVersionResponse struct {

	// 版本信息。
	Versions       *[]Versions `json:"versions,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVersionResponse struct{}"
	}

	return strings.Join([]string{"ShowVersionResponse", string(data)}, " ")
}
