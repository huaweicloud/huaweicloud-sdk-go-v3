package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 配额信息。
type Quota struct {

	// UGO的项目类型。
	ProjectType string `json:"project_type" xml:"project_type"`

	// 总配额。
	Quota int32 `json:"quota" xml:"quota"`

	// 已经使用的配额。
	Used int32 `json:"used" xml:"used"`
}

func (o Quota) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Quota struct{}"
	}

	return strings.Join([]string{"Quota", string(data)}, " ")
}
