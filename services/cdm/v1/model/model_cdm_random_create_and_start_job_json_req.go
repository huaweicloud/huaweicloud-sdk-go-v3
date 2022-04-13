package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CdmRandomCreateAndStartJobJsonReq struct {
	// 作业列表，请参见jobs数据结构说明。

	Jobs []Job `json:"jobs"`
	// CDM集群ID列表，系统会从里面随机选择一个开机状态的集群，在该集群中创建作业并执行作业。

	Clusters []string `json:"clusters"`
}

func (o CdmRandomCreateAndStartJobJsonReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CdmRandomCreateAndStartJobJsonReq struct{}"
	}

	return strings.Join([]string{"CdmRandomCreateAndStartJobJsonReq", string(data)}, " ")
}
