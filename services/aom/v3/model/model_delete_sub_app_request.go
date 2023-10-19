package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSubAppRequest Request Object
type DeleteSubAppRequest struct {

	// 子应用id
	SubAppId string `json:"sub_app_id"`
}

func (o DeleteSubAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSubAppRequest struct{}"
	}

	return strings.Join([]string{"DeleteSubAppRequest", string(data)}, " ")
}
