package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExecutionPlanDiffAttribute struct {

	// 当前资源将要被修改的参数的名字。
	Name *string `json:"name,omitempty"`

	// 当前资源被修改的参数的原始值。  如果是资源创建的场景，则previous_value为空  如果远端资源产生了偏差，则同一个资源会返回两个ExecutionPlanItem，其中一个的drifted为true，另一个的drifted为false   * drifted为true的previous_value为资源栈中所维持的资源属性和状态   * drifted为false的previous_value为provider请求远端资源后，远端资源所返回的资源属性和状态  如果远端资源未产生偏差，则只会返回一个drifted为false的ExecutionPlanItem   * drifted为false的previous_value为资源栈中所维持的资源属性和状态
	PreviousValue *string `json:"previous_value,omitempty"`

	// 当前资源被修改的参数的目的值。  如果是资源删除的场景，则target_value为空  如果远端资源产生了偏差，则同一个资源会返回两个ExecutionPlanItem，其中一个的drifted为true，另一个的drifted为false   * drifted为true的target_value为provider请求远端资源后，远端资源所返回的资源属性和状态   * drifted为false的target_value为基于用户模板更新的资源属性和状态  如果远端资源未产生偏差，则只会返回一个drifted为false的ExecutionPlanItem   * drifted为false的target_value为基于用户模板更新的资源属性和状态
	TargetValue *string `json:"target_value,omitempty"`
}

func (o ExecutionPlanDiffAttribute) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecutionPlanDiffAttribute struct{}"
	}

	return strings.Join([]string{"ExecutionPlanDiffAttribute", string(data)}, " ")
}
