package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RecycleBin struct {

	// 租户project_id
	ProjectId string `json:"project_id"`

	// 回收站配置开关
	Switch string `json:"switch"`

	Policy *RecycleBinPolicys `json:"policy"`
}

func (o RecycleBin) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecycleBin struct{}"
	}

	return strings.Join([]string{"RecycleBin", string(data)}, " ")
}
