package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateGroupResp struct {

	// 创建成功的消费组名称。
	Name *string `json:"name,omitempty"`
}

func (o CreateGroupResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGroupResp struct{}"
	}

	return strings.Join([]string{"CreateGroupResp", string(data)}, " ")
}
