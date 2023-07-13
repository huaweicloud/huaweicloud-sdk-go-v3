package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteApplicationRequest Request Object
type DeleteApplicationRequest struct {

	// 应用ID。
	ApplicationId string `json:"application_id"`
}

func (o DeleteApplicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteApplicationRequest struct{}"
	}

	return strings.Join([]string{"DeleteApplicationRequest", string(data)}, " ")
}
