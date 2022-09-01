package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListResourceGroupRequest struct {

	// 发送的实体的MIME类型。推荐用户默认使用application/json，如果API是对象、镜像上传等接口，媒体类型可按照流类型的不同进行确定。
	ContentType string `json:"Content-Type" xml:"Content-Type"`

	// 资源分组的名称；长度为1-128，只能包含0-9/a-z/A-Z/_/-或汉字；如：ResourceGroup-Test01。
	GroupName *string `json:"group_name,omitempty" xml:"group_name"`

	// 资源分组的ID，长度为1-128，只能包含0-9/a-z/A-Z；如：rg16063743652226ew93e64p。
	GroupId *string `json:"group_id,omitempty" xml:"group_id"`

	// 资源分组健康状态，值可为health、unhealth、no_alarm_rule；health表示健康，
	Status *string `json:"status,omitempty" xml:"status"`

	// 分页起始值，类型为integer，默认值为0。
	Start *int32 `json:"start,omitempty" xml:"start"`

	// 单次查询的条数限制，取值范围(0,100]，默认值为100， 用于限制结果数据条数。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`
}

func (o ListResourceGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceGroupRequest struct{}"
	}

	return strings.Join([]string{"ListResourceGroupRequest", string(data)}, " ")
}
