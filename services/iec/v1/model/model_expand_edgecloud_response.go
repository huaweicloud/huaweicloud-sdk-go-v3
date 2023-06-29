package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandEdgecloudResponse Response Object
type ExpandEdgecloudResponse struct {

	// 边缘业务ID。
	Id *string `json:"id,omitempty"`

	// 部署计划名称。
	Name *string `json:"name,omitempty"`

	// 部署计划ID。
	DeploymentId *string `json:"deployment_id,omitempty"`

	// 边缘业务状态，现存状态有： - creating/scheduling/updating：部署中 - inService：运行中 - failed：创建失败 - deleting：删除中 - delFailed：删除失败
	Status *string `json:"status,omitempty"`

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExpandEdgecloudResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandEdgecloudResponse struct{}"
	}

	return strings.Join([]string{"ExpandEdgecloudResponse", string(data)}, " ")
}
