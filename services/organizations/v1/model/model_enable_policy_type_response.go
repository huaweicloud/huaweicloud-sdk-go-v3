package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type EnablePolicyTypeResponse struct {
	Root           *RootDto `json:"root,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o EnablePolicyTypeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnablePolicyTypeResponse struct{}"
	}

	return strings.Join([]string{"EnablePolicyTypeResponse", string(data)}, " ")
}
