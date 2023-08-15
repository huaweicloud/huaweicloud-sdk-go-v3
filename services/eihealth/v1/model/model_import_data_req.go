package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportDataReq 导入数据请求体
type ImportDataReq struct {

	// 执行策略（true：全部覆盖，false：全部跳过，默认为true）
	Overwrite *bool `json:"overwrite,omitempty"`

	// 源项目ID
	SourceProjectId string `json:"source_project_id"`

	// 导入路径集
	SubPaths []string `json:"sub_paths"`

	// 目标文件夹
	TargetFolder *string `json:"target_folder,omitempty"`
}

func (o ImportDataReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportDataReq struct{}"
	}

	return strings.Join([]string{"ImportDataReq", string(data)}, " ")
}
