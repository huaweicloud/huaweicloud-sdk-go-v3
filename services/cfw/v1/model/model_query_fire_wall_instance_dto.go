package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryFireWallInstanceDto struct {

	// 企业项目id，用户支持企业项目后，由企业项目生成的id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 查询关键字
	KeyWord *string `json:"key_word,omitempty"`

	// 标签列表
	Tags *[]TagInfo `json:"tags,omitempty"`

	// 每页显示个数，范围为1-1024
	Limit int32 `json:"limit"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset int32 `json:"offset"`
}

func (o QueryFireWallInstanceDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryFireWallInstanceDto struct{}"
	}

	return strings.Join([]string{"QueryFireWallInstanceDto", string(data)}, " ")
}
