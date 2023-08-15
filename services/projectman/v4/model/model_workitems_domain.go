package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkitemsDomain id, 领域 14, '性能', 15, '功能', 16, '可靠性' 17, '网络安全' 18, '可维护性' 19, '其他DFX' 20, '可用性' 其他
type WorkitemsDomain struct {

	// 领域id
	Id *string `json:"id,omitempty"`

	// 领域
	Name *string `json:"name,omitempty"`
}

func (o WorkitemsDomain) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkitemsDomain struct{}"
	}

	return strings.Join([]string{"WorkitemsDomain", string(data)}, " ")
}
