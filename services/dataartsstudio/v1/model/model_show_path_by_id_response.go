package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPathByIdResponse Response Object
type ShowPathByIdResponse struct {

	// 路径
	Path           *string `json:"path,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPathByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPathByIdResponse struct{}"
	}

	return strings.Join([]string{"ShowPathByIdResponse", string(data)}, " ")
}
