package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRelatedProjectResponse Response Object
type ShowRelatedProjectResponse struct {
	Result *ShowRelatedProjectResult `json:"result,omitempty"`

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowRelatedProjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRelatedProjectResponse struct{}"
	}

	return strings.Join([]string{"ShowRelatedProjectResponse", string(data)}, " ")
}
