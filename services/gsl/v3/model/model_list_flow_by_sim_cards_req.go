package model

import (
	"encoding/json"

	"strings"
)

type ListFlowBySimCardsReq struct {
	// ICCID列表,最大支持50

	Iccids []string `json:"iccids"`
}

func (o ListFlowBySimCardsReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListFlowBySimCardsReq struct{}"
	}

	return strings.Join([]string{"ListFlowBySimCardsReq", string(data)}, " ")
}
