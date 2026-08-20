package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIpdLabelResponse Response Object
type UpdateIpdLabelResponse struct {

	// 返回状态。
	Status *string `json:"status,omitempty"`

	// 响应信息
	Message *string `json:"message,omitempty"`

	Result         *LabelUpdateResult `json:"result,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o UpdateIpdLabelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpdLabelResponse struct{}"
	}

	return strings.Join([]string{"UpdateIpdLabelResponse", string(data)}, " ")
}
