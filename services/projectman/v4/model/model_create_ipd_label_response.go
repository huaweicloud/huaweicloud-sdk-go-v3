package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIpdLabelResponse Response Object
type CreateIpdLabelResponse struct {

	// 返回状态。
	Status *string `json:"status,omitempty"`

	// 响应信息
	Message *string `json:"message,omitempty"`

	Result         *LabelCreateResult `json:"result,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o CreateIpdLabelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIpdLabelResponse struct{}"
	}

	return strings.Join([]string{"CreateIpdLabelResponse", string(data)}, " ")
}
