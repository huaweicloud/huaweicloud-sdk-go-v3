package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UninstallNextflowRequest Request Object
type UninstallNextflowRequest struct {

	// 引擎ID
	Id string `json:"id"`
}

func (o UninstallNextflowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UninstallNextflowRequest struct{}"
	}

	return strings.Join([]string{"UninstallNextflowRequest", string(data)}, " ")
}
