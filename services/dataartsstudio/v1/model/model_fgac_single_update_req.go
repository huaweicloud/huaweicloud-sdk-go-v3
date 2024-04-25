package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FgacSingleUpdateReq struct {

	// 数据连接id
	DwId *string `json:"dw_id,omitempty"`

	// 是否开启细粒度认证,true表示开启细粒度认证,false表示关闭细粒度认证。
	FgacFlag *bool `json:"fgac_flag,omitempty"`

	// 细粒度认证类型，开启细粒度认证时才生效。\"0\"表示开发态细粒度认证，支持数据开发细粒度脚本运行、作业测试运行，\"1\"表示调度态细粒度认证，支持数据开发细粒度脚本运行、作业测试运行、作业执行调度。
	FgacType *string `json:"fgac_type,omitempty"`
}

func (o FgacSingleUpdateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FgacSingleUpdateReq struct{}"
	}

	return strings.Join([]string{"FgacSingleUpdateReq", string(data)}, " ")
}
