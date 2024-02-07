package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterJobRsp 分子聚类结果
type ClusterJobRsp struct {

	// 分子聚类方法
	Method string `json:"method"`

	// 分子聚类输出结果
	OutputDir string `json:"output_dir"`

	// 作业结果信息
	Status string `json:"status"`

	// 部分失败原因和数量
	FailedReasons []FailedReasonRecord `json:"failed_reasons"`
}

func (o ClusterJobRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterJobRsp struct{}"
	}

	return strings.Join([]string{"ClusterJobRsp", string(data)}, " ")
}
