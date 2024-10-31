package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCaptureResultRequest Request Object
type ListCaptureResultRequest struct {

	// 防火墙id，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取
	FwInstanceId string `json:"fw_instance_id"`

	// 抓包任务id，可通过[查询抓包任务接口](ListCaptureTask.xml)查询获得，通过返回值中的data.records.task_id（.表示各对象之间层级的区分）获得。
	TaskId string `json:"task_id"`

	// 是否指定公网ip下载，0：无限制，1：指定公网ip下载
	Type *int32 `json:"type,omitempty"`

	// 指定公网ip
	Ip *[]string `json:"ip,omitempty"`

	// 企业项目ID，用户根据组织规划企业项目，对应的ID为企业项目ID，可通过[如何获取企业项目ID](cfw_02_0027.xml)获取，用户未开启企业项目时为0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListCaptureResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCaptureResultRequest struct{}"
	}

	return strings.Join([]string{"ListCaptureResultRequest", string(data)}, " ")
}
