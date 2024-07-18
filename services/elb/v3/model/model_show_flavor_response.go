package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlavorResponse Response Object
type ShowFlavorResponse struct {

	// 参数解释：请求ID。  注：自动生成 。
	RequestId *string `json:"request_id,omitempty"`

	Flavor         *Flavor `json:"flavor,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowFlavorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlavorResponse struct{}"
	}

	return strings.Join([]string{"ShowFlavorResponse", string(data)}, " ")
}
