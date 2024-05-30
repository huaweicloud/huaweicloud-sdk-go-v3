package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateProxyNameResponse Response Object
type UpdateProxyNameResponse struct {

	// 修改成功或者失败。
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateProxyNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProxyNameResponse struct{}"
	}

	return strings.Join([]string{"UpdateProxyNameResponse", string(data)}, " ")
}
