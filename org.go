package admin

import (
	`encoding/json`
	`fmt`
	`net/http`

	`github.com/class100/yunke-core`
	`github.com/dgrijalva/jwt-go`
	`github.com/go-resty/resty/v2`
	`github.com/storezhang/gox`
)

// ClientOrg 机构管理客户端
type ClientOrg struct {
	// Url 通信地址
	Url string
	// Name 机构名称
	Name string
	// AuthScheme 加密模式
	AuthScheme string
	// Secret 加密密钥
	Secret string
}

func (co *ClientOrg) requestApi(
	path yunke.ApiPath,
	method gox.HttpMethod,
	params interface{}, pathParams map[string]string,
	version yunke.ApiVersion,
	rsp interface{},
) (err error) {
	return co.request(path, yunke.UrlApiPrefix, method, params, pathParams, version, rsp)
}

func (co *ClientOrg) request(
	path yunke.ApiPath,
	prefix string,
	method gox.HttpMethod,
	params interface{}, pathParams map[string]string,
	version yunke.ApiVersion,
	rsp interface{},
) (err error) {
	var (
		adminRsp           *resty.Response
		authToken          string
		expectedStatusCode int

		url    string
		domain string
	)

	// 修正请求地址为全路径
	orgConfig := yunke.OrgConfig{
		Url:  co.Url,
		Name: co.Name,
	}
	if url, err = orgConfig.GetUrl(path, pathParams, prefix, version); nil != err {
		return
	}
	if domain, err = orgConfig.Domain(); nil != err {
		return
	}

	if authToken, err = token(domain, jwt.SigningMethodHS256, co.Secret); nil != err {
		return
	}

	req := NewResty().SetResult(rsp).SetHeader(gox.HeaderAuthorization, fmt.Sprintf("%s %s", co.AuthScheme, authToken))
	// 注入路径参数
	if 0 != len(pathParams) {
		req = req.SetPathParams(pathParams)
	}

	switch method {
	case gox.HttpMethodGet:
		expectedStatusCode = http.StatusOK

		if nil != params {
			req = req.SetQueryParams(params.(map[string]string))
		}
		adminRsp, err = req.Get(url)
	case gox.HttpMethodPost:
		expectedStatusCode = http.StatusCreated

		if nil != params {
			req = req.SetBody(params)
		}
		adminRsp, err = req.Post(url)
	case gox.HttpMethodPut:
		expectedStatusCode = http.StatusOK

		if nil != params {
			req = req.SetBody(params)
		}
		adminRsp, err = req.Put(url)
	case gox.HttpMethodDelete:
		expectedStatusCode = http.StatusNoContent

		if nil != params {
			req = req.SetBody(params)
		}
		adminRsp, err = req.Delete(url)
	}
	if nil != err {
		return
	}

	if nil == adminRsp {
		err = gox.NewCodeError(gox.ErrorCode(adminRsp.StatusCode()), "无返回数据", RestyStringBody(adminRsp))

		return
	}

	// 检查状态码
	if expectedStatusCode != adminRsp.StatusCode() {
		err = gox.NewCodeError(gox.ErrorCode(adminRsp.StatusCode()), "请求服务器不符合预期", RestyStringBody(adminRsp))
	}

	return
}

func (co ClientOrg) String() string {
	jsonBytes, _ := json.MarshalIndent(co, "", "    ")

	return string(jsonBytes)
}
