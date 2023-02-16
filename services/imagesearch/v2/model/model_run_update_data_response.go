package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RunUpdateDataResponse struct {

	// 更新数据完成返回success。
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunUpdateDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunUpdateDataResponse struct{}"
	}

	return strings.Join([]string{"RunUpdateDataResponse", string(data)}, " ")
}
