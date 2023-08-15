package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnablePolicyTypeResponse Response Object
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
