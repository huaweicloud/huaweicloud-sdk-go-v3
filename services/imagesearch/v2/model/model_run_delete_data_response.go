package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunDeleteDataResponse Response Object
type RunDeleteDataResponse struct {

	// 删除数据完成返回success。
	Result *string `json:"result,omitempty"`

	Data           *DeleteRestInfo `json:"data,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o RunDeleteDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunDeleteDataResponse struct{}"
	}

	return strings.Join([]string{"RunDeleteDataResponse", string(data)}, " ")
}
