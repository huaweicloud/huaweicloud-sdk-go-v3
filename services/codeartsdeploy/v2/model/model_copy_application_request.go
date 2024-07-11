package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CopyApplicationRequest Request Object
type CopyApplicationRequest struct {

	// 应用id
	AppId string `json:"app_id"`
}

func (o CopyApplicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyApplicationRequest struct{}"
	}

	return strings.Join([]string{"CopyApplicationRequest", string(data)}, " ")
}
