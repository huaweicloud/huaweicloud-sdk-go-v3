package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourcesReq struct {

	// 资源信息
	Resources []Resource `json:"resources"`
}

func (o ResourcesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourcesReq struct{}"
	}

	return strings.Join([]string{"ResourcesReq", string(data)}, " ")
}
