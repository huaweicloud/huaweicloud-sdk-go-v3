package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDataProfileResponse Response Object
type ShowDataProfileResponse struct {
	Data *ProfileInfo `json:"data,omitempty"`

	// 行键
	Rowkey *string `json:"rowkey,omitempty"`

	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDataProfileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDataProfileResponse struct{}"
	}

	return strings.Join([]string{"ShowDataProfileResponse", string(data)}, " ")
}
