package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BaseIdResponseData id
type BaseIdResponseData struct {

	// 评估任务id
	Id *string `json:"id,omitempty"`
}

func (o BaseIdResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseIdResponseData struct{}"
	}

	return strings.Join([]string{"BaseIdResponseData", string(data)}, " ")
}
