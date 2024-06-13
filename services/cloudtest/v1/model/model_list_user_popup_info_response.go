package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserPopupInfoResponse Response Object
type ListUserPopupInfoResponse struct {

	// 是否请求成功
	Status *string `json:"status,omitempty"`

	Result *PopUpInfo `json:"result,omitempty"`

	Error          *CommonResponseErrorOfApiTest `json:"error,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListUserPopupInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserPopupInfoResponse struct{}"
	}

	return strings.Join([]string{"ListUserPopupInfoResponse", string(data)}, " ")
}
