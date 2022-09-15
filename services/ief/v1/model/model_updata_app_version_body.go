package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新应用模板版本的请求体
type UpdataAppVersionBody struct {
	Version *VersionUpdate `json:"version"`
}

func (o UpdataAppVersionBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdataAppVersionBody struct{}"
	}

	return strings.Join([]string{"UpdataAppVersionBody", string(data)}, " ")
}
