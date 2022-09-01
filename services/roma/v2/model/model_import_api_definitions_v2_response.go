package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ImportApiDefinitionsV2Response struct {

	// 导入成功信息
	Success *[]Success `json:"success,omitempty" xml:"success"`

	// 导入失败信息
	Failure *[]Failure `json:"failure,omitempty" xml:"failure"`

	Swagger *Swagger `json:"swagger,omitempty" xml:"swagger"`

	// API分组编号
	GroupId *string `json:"group_id,omitempty" xml:"group_id"`

	// 被忽略导入的API信息
	Ignore         *[]Ignore `json:"ignore,omitempty" xml:"ignore"`
	HttpStatusCode int       `json:"-"`
}

func (o ImportApiDefinitionsV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportApiDefinitionsV2Response struct{}"
	}

	return strings.Join([]string{"ImportApiDefinitionsV2Response", string(data)}, " ")
}
