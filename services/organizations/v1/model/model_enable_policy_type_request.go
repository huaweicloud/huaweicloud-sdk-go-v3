package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type EnablePolicyTypeRequest struct {
	Body *PolicyTypeReqBody `json:"body,omitempty"`
}

func (o EnablePolicyTypeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnablePolicyTypeRequest struct{}"
	}

	return strings.Join([]string{"EnablePolicyTypeRequest", string(data)}, " ")
}
