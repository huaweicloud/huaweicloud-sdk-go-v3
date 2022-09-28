package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteEnvResponse struct {

	// 返回ok的结果，表示删除成功
	Ok             *string `json:"ok,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteEnvResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEnvResponse struct{}"
	}

	return strings.Join([]string{"DeleteEnvResponse", string(data)}, " ")
}
