package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnablePolicyAssignmentRequest Request Object
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
