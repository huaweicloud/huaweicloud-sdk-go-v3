package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Events struct {

	// event 批量导入
	Events []Event `json:"events"`
}

func (o Events) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Events struct{}"
	}

	return strings.Join([]string{"Events", string(data)}, " ")
}
