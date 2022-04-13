package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type EnablePolicyAssignmentRequest struct {
	// 规则ID

	PolicyAssignmentId string `json:"policy_assignment_id"`
}

func (o EnablePolicyAssignmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnablePolicyAssignmentRequest struct{}"
	}

	return strings.Join([]string{"EnablePolicyAssignmentRequest", string(data)}, " ")
}
