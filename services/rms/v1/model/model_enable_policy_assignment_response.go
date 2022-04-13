package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type EnablePolicyAssignmentResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o EnablePolicyAssignmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnablePolicyAssignmentResponse struct{}"
	}

	return strings.Join([]string{"EnablePolicyAssignmentResponse", string(data)}, " ")
}
