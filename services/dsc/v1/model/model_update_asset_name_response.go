package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAssetNameResponse Response Object
type UpdateAssetNameResponse struct {

	// 返回消息
	Msg *string `json:"msg,omitempty"`

	// 返回状态，如'200','400'
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateAssetNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAssetNameResponse struct{}"
	}

	return strings.Join([]string{"UpdateAssetNameResponse", string(data)}, " ")
}
