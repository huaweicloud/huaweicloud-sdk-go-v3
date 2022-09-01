package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListPoolsResponse struct {

	// 后端云服务器对象组列表
	Pools          *[]PoolResp `json:"pools,omitempty" xml:"pools"`
	HttpStatusCode int         `json:"-"`
}

func (o ListPoolsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPoolsResponse struct{}"
	}

	return strings.Join([]string{"ListPoolsResponse", string(data)}, " ")
}
