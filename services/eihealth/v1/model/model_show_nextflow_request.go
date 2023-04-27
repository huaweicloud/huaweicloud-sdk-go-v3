package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowNextflowRequest struct {

	// 引擎ID
	Id string `json:"id"`
}

func (o ShowNextflowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNextflowRequest struct{}"
	}

	return strings.Join([]string{"ShowNextflowRequest", string(data)}, " ")
}
