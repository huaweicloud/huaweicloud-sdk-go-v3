package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/smsapi/v1/model"
)

type SMSApiClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewSMSApiClient(hcClient *httpclient.HcHttpClient) *SMSApiClient {
	return &SMSApiClient{HcClient: hcClient}
}

func SMSApiClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder().WithCredentialsType("v1.SMSApiCredentials")
	return builder
}

// BatchSendDiffSms 发送分批短信
//
// 发送分批短信
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SMSApiClient) BatchSendDiffSms(request *model.BatchSendDiffSmsRequest) (*model.BatchSendDiffSmsResponse, error) {
	requestDef := GenReqDefForBatchSendDiffSms()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchSendDiffSmsResponse), nil
	}
}

// BatchSendDiffSmsInvoker 发送分批短信
func (c *SMSApiClient) BatchSendDiffSmsInvoker(request *model.BatchSendDiffSmsRequest) *BatchSendDiffSmsInvoker {
	requestDef := GenReqDefForBatchSendDiffSms()
	return &BatchSendDiffSmsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchSendSms 发送短信
//
// ## 典型场景 ##
//
//	SP调用此接口发送短信。
//
// ## 接口功能 ##
//
//	该接口用于SP请求短信平台向指定用户发送短信。
//
// ## 使用说明 ##
//
//	   前提条件
//	i.  已在短信平台获取该短信能力开放应用的app_key、app_secret。
//
//	  注意事项
//	a.  群发短信时，如果“to”参数携带的号码中包含除数字和+之外的其他字符，则无法向该参数携带的所有号码发送短信。如果“to”参数携带的所有号码只包含数字和+，但部分号码不符合号码规则要求，则在响应消息中会通过状态码标识发送失败的号码，不影响其他正常号码的短信发送。号码之间以英文逗号分隔，每个号码最大长度为21位，最多允许携带500个号码。
//	b. 本接口支持AK/SK认证和X-WSSE认证（二选一），推荐使用AK/SK认证（更安全）。
//	c.  X-Sdk-Date不能与发送请求时的本地时间相差太大（15分钟），否则会导致鉴权失败。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SMSApiClient) BatchSendSms(request *model.BatchSendSmsRequest) (*model.BatchSendSmsResponse, error) {
	requestDef := GenReqDefForBatchSendSms()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchSendSmsResponse), nil
	}
}

// BatchSendSmsInvoker 发送短信
func (c *SMSApiClient) BatchSendSmsInvoker(request *model.BatchSendSmsRequest) *BatchSendSmsInvoker {
	requestDef := GenReqDefForBatchSendSms()
	return &BatchSendSmsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}