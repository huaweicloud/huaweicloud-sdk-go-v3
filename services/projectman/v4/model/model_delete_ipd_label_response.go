package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteIpdLabelResponse Response Object
type DeleteIpdLabelResponse struct {

	// 请求状态
	Status *string `json:"status,omitempty"`

	// 请求失败原因
	Message *string `json:"message,omitempty"`

	Result         *LabelEntity `json:"result,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o DeleteIpdLabelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIpdLabelResponse struct{}"
	}

	return strings.Join([]string{"DeleteIpdLabelResponse", string(data)}, " ")
}
