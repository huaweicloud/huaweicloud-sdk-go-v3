package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddExtensionStarResponse Response Object
type AddExtensionStarResponse struct {

	// 返回值
	Result *interface{} `json:"result,omitempty"`

	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddExtensionStarResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddExtensionStarResponse struct{}"
	}

	return strings.Join([]string{"AddExtensionStarResponse", string(data)}, " ")
}
