package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResizePreparationResponse Response Object
type ShowResizePreparationResponse struct {

	// **参数解释**： 扩容准备的状态。 **取值范围**： INIT：扩容准备初始化中； PREPARING：扩容准备进行中； FAIL：扩容准备失败； SUCCESS：扩容准备成功； WAIT_RETRY：等待扩容准备重试；
	Status *string `json:"status,omitempty"`

	// **参数解释**： 是否支持扩容准备。 **取值范围**： true：支持扩容准备； false：不支持扩容准备；
	IsSupport *bool `json:"is_support,omitempty"`

	// **参数解释**： 扩容准备进度。 **取值范围**： 不涉及
	Progress       *string `json:"progress,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowResizePreparationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResizePreparationResponse struct{}"
	}

	return strings.Join([]string{"ShowResizePreparationResponse", string(data)}, " ")
}
