package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExpackProductInfo struct {

	// 扩展包数量
	ResourceSize *int32 `json:"resource_size,omitempty"`
}

func (o ExpackProductInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpackProductInfo struct{}"
	}

	return strings.Join([]string{"ExpackProductInfo", string(data)}, " ")
}
