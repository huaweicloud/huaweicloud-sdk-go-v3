package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryFireWallInstanceDto struct {

	// 企业项目ID，用户根据组织规划企业项目，对应的ID为企业项目ID，可通过[如何获取企业项目ID](cfw_02_0027.xml)获取，用户未开启企业项目时为0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 查询关键字，可为防火墙id或防火墙名称的一部分。可通过[防火墙ID获取方式](cfw_02_0028.xml)获取
	KeyWord *string `json:"key_word,omitempty"`

	// 标签列表，可通过查询标签服务查询标签接口获得，返回值即为标签列表
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
