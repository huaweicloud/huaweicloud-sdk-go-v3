package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 边缘业务对象
type EdgeCloud struct {
	// 边缘业务ID。

	Id *string `json:"id,omitempty"`
	// 边缘业务名称。 取值范围：只能由中文字符、大小写英文字母、数字及中划线、下划线组成，且长度为[1-32]个字符。

	Name *string `json:"name,omitempty"`
	// 边缘业务描述。最大支持255字节。

	Description *string `json:"description,omitempty"`

	Coverage *CoverageResp `json:"coverage,omitempty"`
	// 创建失败的虚拟机

	FailedNum *int32 `json:"failed_num,omitempty"`
	// 边缘业务状态，现存状态有： - creating/scheduling/updating：部署中 - inService：运行中 - failed：创建失败 - deleting：删除中 - delFailed：删除失败

	Status *string `json:"status,omitempty"`
	// 成功创建的虚拟机

	SuccessNum *int32 `json:"success_num,omitempty"`
	// 边缘业务支持的边缘区域数目。

	EdgeRegions *int32 `json:"edge_regions,omitempty"`
}

func (o EdgeCloud) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EdgeCloud struct{}"
	}

	return strings.Join([]string{"EdgeCloud", string(data)}, " ")
}
