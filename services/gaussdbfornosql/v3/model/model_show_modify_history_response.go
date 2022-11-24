package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowModifyHistoryResponse struct {

	// 实例参数的修改历史列表
	Histories      *[]ConfigurationHistoryRsp `json:"histories,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ShowModifyHistoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowModifyHistoryResponse struct{}"
	}

	return strings.Join([]string{"ShowModifyHistoryResponse", string(data)}, " ")
}
