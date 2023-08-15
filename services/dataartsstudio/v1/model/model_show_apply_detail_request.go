package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowApplyDetailRequest Request Object
type ShowApplyDetailRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// 审核信息id
	ApplyId string `json:"apply_id"`
}

func (o ShowApplyDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApplyDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowApplyDetailRequest", string(data)}, " ")
}
