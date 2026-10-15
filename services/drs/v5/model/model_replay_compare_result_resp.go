package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReplayCompareResultResp 回放库和镜像库里面SQL执行结果不一致详情。
type ReplayCompareResultResp struct {

	// SQL模板
	SqlTemplate *string `json:"sql_template,omitempty"`

	// SQL模板MD5哈希值
	SqlTemplateMd5 *string `json:"sql_template_md5,omitempty"`

	// SQL参数
	SqlParams *string `json:"sql_params,omitempty"`

	// 目标库内容
	TargetContent *string `json:"target_content,omitempty"`

	// 镜像库内容
	MirrorContent *string `json:"mirror_content,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`
}

func (o ReplayCompareResultResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReplayCompareResultResp struct{}"
	}

	return strings.Join([]string{"ReplayCompareResultResp", string(data)}, " ")
}
