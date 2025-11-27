package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComponentUpdateRequest struct {

	// 修改的名称。
	Name string `json:"name"`
}

func (o ComponentUpdateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentUpdateRequest struct{}"
	}

	return strings.Join([]string{"ComponentUpdateRequest", string(data)}, " ")
}
