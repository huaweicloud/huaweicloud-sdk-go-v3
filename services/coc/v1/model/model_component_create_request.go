package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComponentCreateRequest struct {

	// 名称。
	Name string `json:"name"`

	// 应用id。
	ApplicationId string `json:"application_id"`
}

func (o ComponentCreateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentCreateRequest struct{}"
	}

	return strings.Join([]string{"ComponentCreateRequest", string(data)}, " ")
}
