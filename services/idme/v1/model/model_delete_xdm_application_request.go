package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteXdmApplicationRequest Request Object
type DeleteXdmApplicationRequest struct {

	// 应用ID。
	AppId string `json:"app_id"`
}

func (o DeleteXdmApplicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteXdmApplicationRequest struct{}"
	}

	return strings.Join([]string{"DeleteXdmApplicationRequest", string(data)}, " ")
}
