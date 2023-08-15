package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRoutetableRequesBody 更新路由表请求体
type UpdateRoutetableRequesBody struct {
	Routetable *UpdateRoutetableOption `json:"routetable,omitempty"`
}

func (o UpdateRoutetableRequesBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRoutetableRequesBody struct{}"
	}

	return strings.Join([]string{"UpdateRoutetableRequesBody", string(data)}, " ")
}
