package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGraphRequest Request Object
type ShowGraphRequest struct {

	// 图ID。
	GraphId string `json:"graph_id"`
}

func (o ShowGraphRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGraphRequest struct{}"
	}

	return strings.Join([]string{"ShowGraphRequest", string(data)}, " ")
}
