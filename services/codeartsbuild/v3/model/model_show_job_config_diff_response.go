package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobConfigDiffResponse Response Object
type ShowJobConfigDiffResponse struct {

	// 差异内容
	Result *string `json:"result,omitempty"`

	// 状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowJobConfigDiffResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobConfigDiffResponse struct{}"
	}

	return strings.Join([]string{"ShowJobConfigDiffResponse", string(data)}, " ")
}
