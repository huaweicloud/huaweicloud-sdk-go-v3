package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUndoCheckoutByAdminUsingPostResponse Response Object
type ShowUndoCheckoutByAdminUsingPostResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]VersionModelViewDto `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowUndoCheckoutByAdminUsingPostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUndoCheckoutByAdminUsingPostResponse struct{}"
	}

	return strings.Join([]string{"ShowUndoCheckoutByAdminUsingPostResponse", string(data)}, " ")
}
