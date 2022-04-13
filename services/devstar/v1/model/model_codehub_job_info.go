package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CodehubJobInfo struct {
	// 应用名称。

	ApplicationName string `json:"application_name"`
	// 仓库是否私有

	Privately *bool `json:"privately,omitempty"`
	// 仓库短id

	ShortId *string `json:"short_id,omitempty"`
	// 代码存放的ssh地址。

	CodeUrl string `json:"code_url"`
	// CodeHub 仓库所在的 Region ID： - 华南-广州：cn-south-1 - 华东-上海二：cn-east-2 - 华北-北京一：cn-north-1 - 华北-北京四：cn-north-4

	RegionId string `json:"region_id"`
	// - 0 - 将生成的应用代码存储于 repo_info 指定的 CodeHub 仓库中。 - 1 - 将生成的应用代码存储到华为云，任务创建人可以通过 ExportApplicationCode 下载代码压缩包。

	RepoType int32 `json:"repo_type"`
	// 可以根据 template-metadata.json 获取动态参数 ID 以及规则。

	Properties map[string]string `json:"properties,omitempty"`

	RepoInfo *RepositoryInfo `json:"repo_info,omitempty"`
}

func (o CodehubJobInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CodehubJobInfo struct{}"
	}

	return strings.Join([]string{"CodehubJobInfo", string(data)}, " ")
}
