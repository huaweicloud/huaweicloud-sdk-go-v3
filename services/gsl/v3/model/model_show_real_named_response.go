package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowRealNamedResponse struct {

	// ICCID
	Iccid *string `json:"iccid,omitempty" xml:"iccid"`

	// 是否已实名认证: true表示是，false表示否。
	RealNamed      *bool `json:"real_named,omitempty" xml:"real_named"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowRealNamedResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRealNamedResponse struct{}"
	}

	return strings.Join([]string{"ShowRealNamedResponse", string(data)}, " ")
}
