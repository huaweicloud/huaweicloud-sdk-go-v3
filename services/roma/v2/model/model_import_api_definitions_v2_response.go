package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ImportApiDefinitionsV2Response struct {

	// 导入成功信息
	Success *[]Success `json:"success,omitempty"`

	// 导入失败信息
	Failure *[]Failure `json:"failure,omitempty"`

	Swagger *Swagger `json:"swagger,omitempty"`

	// API分组编号
	GroupId *string `json:"group_id,omitempty"`

	// 被忽略导入的API信息
	Ignore         *[]Ignore `json:"ignore,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ImportApiDefinitionsV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportApiDefinitionsV2Response struct{}"
	}

	return strings.Join([]string{"ImportApiDefinitionsV2Response", string(data)}, " ")
}
