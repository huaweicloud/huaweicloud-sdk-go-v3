package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowApplyHistoryResponse Response Object
type ShowApplyHistoryResponse struct {

	// 参数组模板应用历史列表
	Histories      *[]ApplyHistoryRsp `json:"histories,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowApplyHistoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApplyHistoryResponse struct{}"
	}

	return strings.Join([]string{"ShowApplyHistoryResponse", string(data)}, " ")
}
