package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CloneDataReq 复制数据请求体
type CloneDataReq struct {

	// 执行策略（true：全部覆盖，false：全部跳过，默认为true）
	Overwrite *bool `json:"overwrite,omitempty"`

	// 复制的路径集
	SubPaths []string `json:"sub_paths"`

	// 目标文件夹
	TargetFolder *string `json:"target_folder,omitempty"`
}

func (o CloneDataReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloneDataReq struct{}"
	}

	return strings.Join([]string{"CloneDataReq", string(data)}, " ")
}
