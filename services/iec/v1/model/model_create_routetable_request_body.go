package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建路由表请求体
type CreateRoutetableRequestBody struct {
	Routetable *CreateRoutetableOption `json:"routetable"`
}

func (o CreateRoutetableRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRoutetableRequestBody struct{}"
	}

	return strings.Join([]string{"CreateRoutetableRequestBody", string(data)}, " ")
}
