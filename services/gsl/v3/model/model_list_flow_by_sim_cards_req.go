package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListFlowBySimCardsReq struct {
	// ICCID列表,最大支持50

	Iccids []string `json:"iccids"`
}

func (o ListFlowBySimCardsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlowBySimCardsReq struct{}"
	}

	return strings.Join([]string{"ListFlowBySimCardsReq", string(data)}, " ")
}
