/*
 * RMS
 *
 * Resource Manager Api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type EnablePolicyAssignmentResponse struct {
}

func (o EnablePolicyAssignmentResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"EnablePolicyAssignmentResponse", string(data)}, " ")
}
