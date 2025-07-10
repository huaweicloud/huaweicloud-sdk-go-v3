package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTerminalsBindingDesktopsResponse Response Object
type DeleteTerminalsBindingDesktopsResponse struct {

	// 需删除的策略ID列表。
	ResultList     *[]DeleteTerminalsBindingDesktopsResult `json:"result_list,omitempty"`
	HttpStatusCode int                                     `json:"-"`
}

func (o DeleteTerminalsBindingDesktopsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTerminalsBindingDesktopsResponse struct{}"
	}

	return strings.Join([]string{"DeleteTerminalsBindingDesktopsResponse", string(data)}, " ")
}
