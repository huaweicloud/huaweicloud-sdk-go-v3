package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ORACLE集群地址
type MultiOracleAddress struct {

	// ORACLE地址
	OracleAddress *string `json:"oracle_address,omitempty" xml:"oracle_address"`

	// ORACLE端口
	OraclePort *string `json:"oracle_port,omitempty" xml:"oracle_port"`
}

func (o MultiOracleAddress) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiOracleAddress struct{}"
	}

	return strings.Join([]string{"MultiOracleAddress", string(data)}, " ")
}
