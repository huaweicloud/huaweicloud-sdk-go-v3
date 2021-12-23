package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 输出为localpath类型时的配置信息
type TaskOutputLocalpath struct {
	// 挂载的源路径，存放作业运行结果的路径，选用localpath类型输出时必填。必须为linux路径，且为选用的边缘运行池挂载路径的子路径，长度不超过255。

	MountSourcePath string `json:"mount_source_path"`
	// 作业输出数据类别的列表，选填，仅部分服务需要。当输出类型下有这个列表时，表示希望这个输出类型下存放dataCategory列表内的数据。

	DataCategory *[]string `json:"data_category,omitempty"`
}

func (o TaskOutputLocalpath) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskOutputLocalpath struct{}"
	}

	return strings.Join([]string{"TaskOutputLocalpath", string(data)}, " ")
}
