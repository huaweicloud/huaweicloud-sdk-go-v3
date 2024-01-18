package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeServerImageRequest Request Object
type ChangeServerImageRequest struct {

	// 服务器唯一标识。
	ServerId string `json:"server_id"`

	Body *ChangeServerImageReq `json:"body,omitempty"`
}

func (o ChangeServerImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeServerImageRequest struct{}"
	}

	return strings.Join([]string{"ChangeServerImageRequest", string(data)}, " ")
}
