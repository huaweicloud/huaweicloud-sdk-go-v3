package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAppBody 更新应用模板的请求体
type UpdateAppBody struct {
	App *AppUpdate `json:"app"`
}

func (o UpdateAppBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppBody struct{}"
	}

	return strings.Join([]string{"UpdateAppBody", string(data)}, " ")
}
