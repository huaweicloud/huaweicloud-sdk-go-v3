package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ccm/v1/model"
)

type CcmClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCcmClient(hcClient *http_client.HcHttpClient) *CcmClient {
	return &CcmClient{HcClient: hcClient}
}

func CcmClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

// CreateCertificate 申请证书
//
// 申请证书。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CcmClient) CreateCertificate(request *model.CreateCertificateRequest) (*model.CreateCertificateResponse, error) {
	requestDef := GenReqDefForCreateCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCertificateResponse), nil
	}
}

// CreateCertificateInvoker 申请证书
func (c *CcmClient) CreateCertificateInvoker(request *model.CreateCertificateRequest) *CreateCertificateInvoker {
	requestDef := GenReqDefForCreateCertificate()
	return &CreateCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCertificateAuthorityObsAgency 创建委托
//
// 用户给PCA创建OBS委托授权，用于访问OBS桶，更新吊销列表。
// &gt; 用户所使用账号token需要具备安全管理员（secu_admin）权限。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CcmClient) CreateCertificateAuthorityObsAgency(request *model.CreateCertificateAuthorityObsAgencyRequest) (*model.CreateCertificateAuthorityObsAgencyResponse, error) {
	requestDef := GenReqDefForCreateCertificateAuthorityObsAgency()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCertificateAuthorityObsAgencyResponse), nil
	}
}

// CreateCertificateAuthorityObsAgencyInvoker 创建委托
func (c *CcmClient) CreateCertificateAuthorityObsAgencyInvoker(request *model.CreateCertificateAuthorityObsAgencyRequest) *CreateCertificateAuthorityObsAgencyInvoker {
	requestDef := GenReqDefForCreateCertificateAuthorityObsAgency()
	return &CreateCertificateAuthorityObsAgencyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCertificateByCsr 通过CSR签发证书
//
// 通过CSR签发证书。功能约束如下：
// - 1、当前默认参数如下：
//   - CA 默认参数：
//       - **keyUsage**: digitalSignature, keyCertSign, cRLSign，优先采用CSR中的参数；
//       - **SignatureHashAlgorithm**: SHA384；
//       - **PathLength**：0 （可自定义）。
//   - 私有证书：
//       - **keyUsage**: digitalSignature keyAgreement，优先采用CSR中的参数；
//       - **SignatureHashAlgorithm**: SHA384；
// - 2、当传入的type为**INTERMEDIATE_CA**时，创建出的从属CA证书，有以下限制：
//   - 不占用CA配额。在查询CA列表时，不会返回该证书；
//   - 只支持通过以下两个接口获取其信息：
//       - GET /v1/private-certificate-authorities/{ca_id} 获取证书详情
//       - POST /v1/private-certificate-authorities/{ca_id}/export 导出证书
//   - 本接口返回的**certificate_id**即代表从属CA的**ca_id**；
//   - 无法用于签发证书，密钥在用户侧。
// - 3、当type为**ENTITY_CERT**时，创建出的私有证书，有以下特点：
//   - 占用私有证书配额。在查询私有证书列表时，会返回该证书；
//   - 除了导出时不包含密钥信息（密钥在用户端），其余用法与其它私有证书一致。
// &gt; 注：需要使用“\\r\\n”或“\\n”代替换行符，将CSR转换成一串字符，可参考示例请求。注：目前，证书的组织信息、公钥算法以及公钥内容等均来自CSR文件，暂不支持API传入。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CcmClient) CreateCertificateByCsr(request *model.CreateCertificateByCsrRequest) (*model.CreateCertificateByCsrResponse, error) {
	requestDef := GenReqDefForCreateCertificateByCsr()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCertificateByCsrResponse), nil
	}
}

// CreateCertificateByCsrInvoker 通过CSR签发证书
func (c *CcmClient) CreateCertificateByCsrInvoker(request *model.CreateCertificateByCsrRequest) *CreateCertificateByCsrInvoker {
	requestDef := GenReqDefForCreateCertificateByCsr()
	return &CreateCertificateByCsrInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCertificate 删除证书
//
// 删除证书。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CcmClient) DeleteCertificate(request *model.DeleteCertificateRequest) (*model.DeleteCertificateResponse, error) {
	requestDef := GenReqDefForDeleteCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCertificateResponse), nil
	}
}

// DeleteCertificateInvoker 删除证书
func (c *CcmClient) DeleteCertificateInvoker(request *model.DeleteCertificateRequest) *DeleteCertificateInvoker {
	requestDef := GenReqDefForDeleteCertificate()
	return &DeleteCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportCertificate 导出证书
//
// 导出证书。
//   - 选择是否压缩时，分以下两种情况：
//     - is_compressed为true时，返回文件压缩包，命名为：证书名称_type字段小写字母.zip，如”test_apache.zip“。
//       - type &#x3D; \&quot;**APACHE**\&quot;时，压缩包中包含三个文件：**server.key**（密钥文件，内容为PEM格式）、**chain.crt**（证书链，内容为PEM格式）、**server.crt**（证书，内容为PEM格式）；
//       - type &#x3D; \&quot;**IIS**\&quot;时，压缩包中包含两个文件：**keystorePass.txt**（keystore口令）、**server.pfx**（PFX证书，证书与证书链包含在同一个文件）；
//       - type &#x3D; \&quot;**NGINX**\&quot;时，压缩包中包含两个文件：**server.key**（密钥文件，内容为PEM格式）、**server.crt**（内容为PEM格式，证书与证书链包含在同一个文件）；
//       - type &#x3D; \&quot;**TOMCAT**\&quot;时，压缩包中包含两个文件：**keystorePass.txt**（keystore口令）、**server.jks**（JKX证书，证书与证书链包含在同一个文件）；
//       - type &#x3D; \&quot;**OTHER**\&quot;时，压缩包中包含三个文件：**server.key**（密钥文件，内容为PEM格式）、**chain.pem**（证书链）、**server.pem**（证书）。
//     - is_compressed为false时，返回json格式，返回的具体参数如下：
//       - type &#x3D; \&quot;**APACHE**\&quot;或\&quot;**NGINX**\&quot;或\&quot;**OTHER**\&quot;时，返回参数如下：
//         - **certificate**（证书内容，PEM格式）；
//         - **certificate_chain**（证书链，PEM格式）；
//         - **private_key**（证书私钥，PEM格式）；
//       - type &#x3D; \&quot;**IIS**\&quot;或\&quot;**TOMCAT**\&quot;时，暂时未定义。
// &gt; 只有当证书状态为“已签发”时，可进行导出操作。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CcmClient) ExportCertificate(request *model.ExportCertificateRequest) (*model.ExportCertificateResponse, error) {
	requestDef := GenReqDefForExportCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportCertificateResponse), nil
	}
}

// ExportCertificateInvoker 导出证书
func (c *CcmClient) ExportCertificateInvoker(request *model.ExportCertificateRequest) *ExportCertificateInvoker {
	requestDef := GenReqDefForExportCertificate()
	return &ExportCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCertificate 查询私有证书列表
//
// 查询私有证书列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CcmClient) ListCertificate(request *model.ListCertificateRequest) (*model.ListCertificateResponse, error) {
	requestDef := GenReqDefForListCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCertificateResponse), nil
	}
}

// ListCertificateInvoker 查询私有证书列表
func (c *CcmClient) ListCertificateInvoker(request *model.ListCertificateRequest) *ListCertificateInvoker {
	requestDef := GenReqDefForListCertificate()
	return &ListCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCertificateAuthorityObsBucket 查询OBS桶列表
//
// 查询OBS桶列表。
// &gt; 只有用户创建了委托授权，方可使用此接口。创建委托授权参见本文档：**证书吊销处理&gt;创建委托**。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CcmClient) ListCertificateAuthorityObsBucket(request *model.ListCertificateAuthorityObsBucketRequest) (*model.ListCertificateAuthorityObsBucketResponse, error) {
	requestDef := GenReqDefForListCertificateAuthorityObsBucket()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCertificateAuthorityObsBucketResponse), nil
	}
}

// ListCertificateAuthorityObsBucketInvoker 查询OBS桶列表
func (c *CcmClient) ListCertificateAuthorityObsBucketInvoker(request *model.ListCertificateAuthorityObsBucketRequest) *ListCertificateAuthorityObsBucketInvoker {
	requestDef := GenReqDefForListCertificateAuthorityObsBucket()
	return &ListCertificateAuthorityObsBucketInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ParseCertificateSigningRequest 解析CSR
//
// 解析CSR。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CcmClient) ParseCertificateSigningRequest(request *model.ParseCertificateSigningRequestRequest) (*model.ParseCertificateSigningRequestResponse, error) {
	requestDef := GenReqDefForParseCertificateSigningRequest()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ParseCertificateSigningRequestResponse), nil
	}
}

// ParseCertificateSigningRequestInvoker 解析CSR
func (c *CcmClient) ParseCertificateSigningRequestInvoker(request *model.ParseCertificateSigningRequestRequest) *ParseCertificateSigningRequestInvoker {
	requestDef := GenReqDefForParseCertificateSigningRequest()
	return &ParseCertificateSigningRequestInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RevokeCertificate 吊销证书
//
// 吊销证书。
// &gt; 注：当不想填写吊销理由时，请求body体请置为\&quot;**{}**\&quot;，否则将会报错。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CcmClient) RevokeCertificate(request *model.RevokeCertificateRequest) (*model.RevokeCertificateResponse, error) {
	requestDef := GenReqDefForRevokeCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RevokeCertificateResponse), nil
	}
}

// RevokeCertificateInvoker 吊销证书
func (c *CcmClient) RevokeCertificateInvoker(request *model.RevokeCertificateRequest) *RevokeCertificateInvoker {
	requestDef := GenReqDefForRevokeCertificate()
	return &RevokeCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCertificate 查询证书详情
//
// 查询证书详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CcmClient) ShowCertificate(request *model.ShowCertificateRequest) (*model.ShowCertificateResponse, error) {
	requestDef := GenReqDefForShowCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCertificateResponse), nil
	}
}

// ShowCertificateInvoker 查询证书详情
func (c *CcmClient) ShowCertificateInvoker(request *model.ShowCertificateRequest) *ShowCertificateInvoker {
	requestDef := GenReqDefForShowCertificate()
	return &ShowCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCertificateAuthorityObsAgency 查看是否具有委托权限
//
// 查看是否具有委托权限。
// &gt; 用户所使用账号token需要具备安全管理员（secu_admin）权限。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CcmClient) ShowCertificateAuthorityObsAgency(request *model.ShowCertificateAuthorityObsAgencyRequest) (*model.ShowCertificateAuthorityObsAgencyResponse, error) {
	requestDef := GenReqDefForShowCertificateAuthorityObsAgency()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCertificateAuthorityObsAgencyResponse), nil
	}
}

// ShowCertificateAuthorityObsAgencyInvoker 查看是否具有委托权限
func (c *CcmClient) ShowCertificateAuthorityObsAgencyInvoker(request *model.ShowCertificateAuthorityObsAgencyRequest) *ShowCertificateAuthorityObsAgencyInvoker {
	requestDef := GenReqDefForShowCertificateAuthorityObsAgency()
	return &ShowCertificateAuthorityObsAgencyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCertificateQuota 查询私有证书配额
//
// 查询私有证书配额。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CcmClient) ShowCertificateQuota(request *model.ShowCertificateQuotaRequest) (*model.ShowCertificateQuotaResponse, error) {
	requestDef := GenReqDefForShowCertificateQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCertificateQuotaResponse), nil
	}
}

// ShowCertificateQuotaInvoker 查询私有证书配额
func (c *CcmClient) ShowCertificateQuotaInvoker(request *model.ShowCertificateQuotaRequest) *ShowCertificateQuotaInvoker {
	requestDef := GenReqDefForShowCertificateQuota()
	return &ShowCertificateQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCertificateAuthority 创建CA
//
// 创建CA，分以下三种情况：
// - 创建根CA，根据参数介绍中，填写必选值；
// - 创建从属CA，并需要直接激活该证书，根据参数介绍中，填写必选值；
// - 创建从属CA，不需要直接激活该证书，请求body中只需要缺少此三个参数之一即可：issuer_id、signature_algorithm、validity。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CcmClient) CreateCertificateAuthority(request *model.CreateCertificateAuthorityRequest) (*model.CreateCertificateAuthorityResponse, error) {
	requestDef := GenReqDefForCreateCertificateAuthority()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCertificateAuthorityResponse), nil
	}
}

// CreateCertificateAuthorityInvoker 创建CA
func (c *CcmClient) CreateCertificateAuthorityInvoker(request *model.CreateCertificateAuthorityRequest) *CreateCertificateAuthorityInvoker {
	requestDef := GenReqDefForCreateCertificateAuthority()
	return &CreateCertificateAuthorityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCertificateAuthority 删除CA
//
// 计划删除CA。计划多少天后删除CA证书，可设置7天～30天内删除。
// &gt; 只有当证书状态为”待激活“或”已禁用“状态时，才可删除。”待激活“状态下，将会立即删除证书，不支持延迟删除。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CcmClient) DeleteCertificateAuthority(request *model.DeleteCertificateAuthorityRequest) (*model.DeleteCertificateAuthorityResponse, error) {
	requestDef := GenReqDefForDeleteCertificateAuthority()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCertificateAuthorityResponse), nil
	}
}

// DeleteCertificateAuthorityInvoker 删除CA
func (c *CcmClient) DeleteCertificateAuthorityInvoker(request *model.DeleteCertificateAuthorityRequest) *DeleteCertificateAuthorityInvoker {
	requestDef := GenReqDefForDeleteCertificateAuthority()
	return &DeleteCertificateAuthorityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisableCertificateAuthority 禁用CA
//
// 禁用CA。
// &gt; 只有当证书处于\&quot;已激活\&quot;或\&quot;已过期\&quot;状态时，可进行禁用操作。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CcmClient) DisableCertificateAuthority(request *model.DisableCertificateAuthorityRequest) (*model.DisableCertificateAuthorityResponse, error) {
	requestDef := GenReqDefForDisableCertificateAuthority()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableCertificateAuthorityResponse), nil
	}
}

// DisableCertificateAuthorityInvoker 禁用CA
func (c *CcmClient) DisableCertificateAuthorityInvoker(request *model.DisableCertificateAuthorityRequest) *DisableCertificateAuthorityInvoker {
	requestDef := GenReqDefForDisableCertificateAuthority()
	return &DisableCertificateAuthorityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableCertificateAuthority 启用CA
//
// 启用CA。
// &gt; 注：只有当证书处于\&quot;已禁用\&quot;状态时，可进行启用操作。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CcmClient) EnableCertificateAuthority(request *model.EnableCertificateAuthorityRequest) (*model.EnableCertificateAuthorityResponse, error) {
	requestDef := GenReqDefForEnableCertificateAuthority()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableCertificateAuthorityResponse), nil
	}
}

// EnableCertificateAuthorityInvoker 启用CA
func (c *CcmClient) EnableCertificateAuthorityInvoker(request *model.EnableCertificateAuthorityRequest) *EnableCertificateAuthorityInvoker {
	requestDef := GenReqDefForEnableCertificateAuthority()
	return &EnableCertificateAuthorityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportCertificateAuthorityCertificate 导出CA证书
//
// 导出CA证书。
// &gt; 注：只有当证书处于\&quot;已激活\&quot;或\&quot;已过期\&quot;时，可进行导出操作。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CcmClient) ExportCertificateAuthorityCertificate(request *model.ExportCertificateAuthorityCertificateRequest) (*model.ExportCertificateAuthorityCertificateResponse, error) {
	requestDef := GenReqDefForExportCertificateAuthorityCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportCertificateAuthorityCertificateResponse), nil
	}
}

// ExportCertificateAuthorityCertificateInvoker 导出CA证书
func (c *CcmClient) ExportCertificateAuthorityCertificateInvoker(request *model.ExportCertificateAuthorityCertificateRequest) *ExportCertificateAuthorityCertificateInvoker {
	requestDef := GenReqDefForExportCertificateAuthorityCertificate()
	return &ExportCertificateAuthorityCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportCertificateAuthorityCsr 导出CA的证书签名请求（CSR）
//
// 导出CA的证书签名请求。
// &gt; 只有当CA处于\&quot;待激活\&quot;状态时，可导出证书签名请求。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CcmClient) ExportCertificateAuthorityCsr(request *model.ExportCertificateAuthorityCsrRequest) (*model.ExportCertificateAuthorityCsrResponse, error) {
	requestDef := GenReqDefForExportCertificateAuthorityCsr()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportCertificateAuthorityCsrResponse), nil
	}
}

// ExportCertificateAuthorityCsrInvoker 导出CA的证书签名请求（CSR）
func (c *CcmClient) ExportCertificateAuthorityCsrInvoker(request *model.ExportCertificateAuthorityCsrRequest) *ExportCertificateAuthorityCsrInvoker {
	requestDef := GenReqDefForExportCertificateAuthorityCsr()
	return &ExportCertificateAuthorityCsrInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportCertificateAuthorityCertificate 导入CA证书
//
// 导入CA证书，使用本接口需要满足以下条件：
//   - （1）证书为“待激活”状态的从属CA；
//   - （2）导入的证书体必须满足以下条件：
//       - a、该证书被签发时的证书签名请求必须是从PCA系统中导出；
//       - b、其证书链虽然允许不上传，但后期若想要导出完整的证书链，应导入完整的证书链；
//       - c、证书体与证书链必须为PEM编码。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CcmClient) ImportCertificateAuthorityCertificate(request *model.ImportCertificateAuthorityCertificateRequest) (*model.ImportCertificateAuthorityCertificateResponse, error) {
	requestDef := GenReqDefForImportCertificateAuthorityCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportCertificateAuthorityCertificateResponse), nil
	}
}

// ImportCertificateAuthorityCertificateInvoker 导入CA证书
func (c *CcmClient) ImportCertificateAuthorityCertificateInvoker(request *model.ImportCertificateAuthorityCertificateRequest) *ImportCertificateAuthorityCertificateInvoker {
	requestDef := GenReqDefForImportCertificateAuthorityCertificate()
	return &ImportCertificateAuthorityCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// IssueCertificateAuthorityCertificate 激活CA
//
// 激活CA。
// &gt; 只有当证书处于\&quot;待激活\&quot;状态时，可进行激活操作。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CcmClient) IssueCertificateAuthorityCertificate(request *model.IssueCertificateAuthorityCertificateRequest) (*model.IssueCertificateAuthorityCertificateResponse, error) {
	requestDef := GenReqDefForIssueCertificateAuthorityCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.IssueCertificateAuthorityCertificateResponse), nil
	}
}

// IssueCertificateAuthorityCertificateInvoker 激活CA
func (c *CcmClient) IssueCertificateAuthorityCertificateInvoker(request *model.IssueCertificateAuthorityCertificateRequest) *IssueCertificateAuthorityCertificateInvoker {
	requestDef := GenReqDefForIssueCertificateAuthorityCertificate()
	return &IssueCertificateAuthorityCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCertificateAuthority 查询CA列表
//
// 查询CA列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CcmClient) ListCertificateAuthority(request *model.ListCertificateAuthorityRequest) (*model.ListCertificateAuthorityResponse, error) {
	requestDef := GenReqDefForListCertificateAuthority()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCertificateAuthorityResponse), nil
	}
}

// ListCertificateAuthorityInvoker 查询CA列表
func (c *CcmClient) ListCertificateAuthorityInvoker(request *model.ListCertificateAuthorityRequest) *ListCertificateAuthorityInvoker {
	requestDef := GenReqDefForListCertificateAuthority()
	return &ListCertificateAuthorityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestoreCertificateAuthority 恢复CA
//
// 恢复CA，将处于“计划删除”状态的CA证书，重新恢复为“已禁用”状态。
// &gt; 注：只有处于“计划删除”状态的CA证书，才可进行恢复操作。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CcmClient) RestoreCertificateAuthority(request *model.RestoreCertificateAuthorityRequest) (*model.RestoreCertificateAuthorityResponse, error) {
	requestDef := GenReqDefForRestoreCertificateAuthority()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreCertificateAuthorityResponse), nil
	}
}

// RestoreCertificateAuthorityInvoker 恢复CA
func (c *CcmClient) RestoreCertificateAuthorityInvoker(request *model.RestoreCertificateAuthorityRequest) *RestoreCertificateAuthorityInvoker {
	requestDef := GenReqDefForRestoreCertificateAuthority()
	return &RestoreCertificateAuthorityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RevokeCertificateAuthority 吊销CA
//
// 吊销子CA。
// &gt; 注：当不想填写吊销理由时，请求body体请置为\&quot;**{}**\&quot;，否则将会报错。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CcmClient) RevokeCertificateAuthority(request *model.RevokeCertificateAuthorityRequest) (*model.RevokeCertificateAuthorityResponse, error) {
	requestDef := GenReqDefForRevokeCertificateAuthority()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RevokeCertificateAuthorityResponse), nil
	}
}

// RevokeCertificateAuthorityInvoker 吊销CA
func (c *CcmClient) RevokeCertificateAuthorityInvoker(request *model.RevokeCertificateAuthorityRequest) *RevokeCertificateAuthorityInvoker {
	requestDef := GenReqDefForRevokeCertificateAuthority()
	return &RevokeCertificateAuthorityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCertificateAuthority 查询CA详情
//
// 查询CA详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CcmClient) ShowCertificateAuthority(request *model.ShowCertificateAuthorityRequest) (*model.ShowCertificateAuthorityResponse, error) {
	requestDef := GenReqDefForShowCertificateAuthority()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCertificateAuthorityResponse), nil
	}
}

// ShowCertificateAuthorityInvoker 查询CA详情
func (c *CcmClient) ShowCertificateAuthorityInvoker(request *model.ShowCertificateAuthorityRequest) *ShowCertificateAuthorityInvoker {
	requestDef := GenReqDefForShowCertificateAuthority()
	return &ShowCertificateAuthorityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCertificateAuthorityQuota 查询CA配额
//
// 查询CA证书配额。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CcmClient) ShowCertificateAuthorityQuota(request *model.ShowCertificateAuthorityQuotaRequest) (*model.ShowCertificateAuthorityQuotaResponse, error) {
	requestDef := GenReqDefForShowCertificateAuthorityQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCertificateAuthorityQuotaResponse), nil
	}
}

// ShowCertificateAuthorityQuotaInvoker 查询CA配额
func (c *CcmClient) ShowCertificateAuthorityQuotaInvoker(request *model.ShowCertificateAuthorityQuotaRequest) *ShowCertificateAuthorityQuotaInvoker {
	requestDef := GenReqDefForShowCertificateAuthorityQuota()
	return &ShowCertificateAuthorityQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
