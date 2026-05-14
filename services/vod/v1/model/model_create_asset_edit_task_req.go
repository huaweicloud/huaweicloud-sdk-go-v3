package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateAssetEditTaskReq struct {

	// 待编辑媒资列表，最多支持20个。
	Inputs []EditInput `json:"inputs"`

	EditingSettings *VodEditingSetting `json:"editing_settings,omitempty"`
}

func (o CreateAssetEditTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAssetEditTaskReq struct{}"
	}

	return strings.Join([]string{"CreateAssetEditTaskReq", string(data)}, " ")
}
