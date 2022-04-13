package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ORACLE集群地址
type MultiOracleAddress struct {
	// ORACLE地址

	OracleAddress *string `json:"oracle_address,omitempty"`
	// ORACLE端口

	OraclePort *string `json:"oracle_port,omitempty"`
}

func (o MultiOracleAddress) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiOracleAddress struct{}"
	}

	return strings.Join([]string{"MultiOracleAddress", string(data)}, " ")
}
