package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClusteringJobReq 创建聚类分析作业请求体。
type CreateClusteringJobReq struct {
	BasicInfo *CreateDrugJobBasicInfo `json:"basic_info"`

	File *ClusteringDrugFile `json:"file"`
}

func (o CreateClusteringJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusteringJobReq struct{}"
	}

	return strings.Join([]string{"CreateClusteringJobReq", string(data)}, " ")
}
