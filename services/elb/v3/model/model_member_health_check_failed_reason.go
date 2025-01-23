package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MemberHealthCheckFailedReason 健康检查异常的原因。
type MemberHealthCheckFailedReason struct {

	// 参数解释：健康检查异常原因码。 取值范围： - CONNECT_TIMEOUT: 负载均衡健康检查时向后端服务器建立连接超时。 - CONNECT_REFUSED: 负载均衡健康检查时向后端服务器拒绝连接。 - CONNECT_FAILED: 负载均衡健康检查时向后端服务器建立连接失败。 - CONNECT_INTERRUPT: 负载均衡健康检查与后端服务器连接中断。 - SSL_HANDSHAKE_ERROR: 负载均衡健康检查时与后端服务器SSL握手失败。 - RECV_RESPONSE_FAILED: 负载均衡健康检查时从后端服务器接收响应失败。 - RECV_RESPONSE_TIMEOUT: 负载均衡健康检查时从后端服务器接收响应超时。 - SEND_REQUEST_FAILED: 负载均衡健康检查时向后端服务器发送请求失败。 - SEND_REQUEST_TIMEOUT: 负载均衡健康检查时向后端服务器发送请求超时。 - RESPONSE_FORMAT_ERROR: 负载均衡健康检查时从后端服务器接收响应格式错误。 - RESPONSE_MISMATCH: 负载均衡健康检查时从后端服务器接收的响应码与预期配置返回码不一致。
	ReasonCode string `json:"reason_code"`

	// 参数解释：健康检查预期响应状态码。 支持HTTP,HTTPS,GRPC健康检查。 只有reason_code为RESPONSE_MISMATCH时，支持非null取值。  取值范围： - 单值：单个返回码。当type为GRPC时，取值范围为0-99；当type为其他协议是，取值范围为200-599。例如：\"0\"或\"200\"。 - 列表：多个特定返回码，逗号分隔，最多支持5个返回码。例如:\"200,202\"或\"0,1\"。 - 区间：一个返回码区间，区间内\"-\"分隔，区间之间逗号分隔，最多支持5个区间。例如\"200-204,300-399\"或\"0-5,10-12,20-30\"。
	ExpectedResponse string `json:"expected_response"`

	// 参数解释：健康检查探测实际响应状态码。 支持HTTP,HTTPS,GRPC健康检查。 只有reason_code为RESPONSE_MISMATCH时，支持非null取值。  取值范围： - 单个返回码。当type为GRPC时，取值范围为0-99；当type为其他协议时，取值范围为200-599。例如：\"0\"或\"200\"。
	HealthcheckResponse string `json:"healthcheck_response"`
}

func (o MemberHealthCheckFailedReason) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MemberHealthCheckFailedReason struct{}"
	}

	return strings.Join([]string{"MemberHealthCheckFailedReason", string(data)}, " ")
}
