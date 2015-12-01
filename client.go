package xinge
import(
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

type Client struct {
	accessId  int
	secretKey string
}
//推送给单个设备
func (c *Client) PushSingleDevice(deviceType int,deviceToken string, message CommonMessage) Response {
	params := make(map[string]interface{})
	params["device_token"] = deviceToken
	params["message"] = string(message.XGjson())
	res := c.push(deviceType,RESTAPI_PUSHSINGLEDEVICE,params)
	return res
}
//推送给单个账户或别名
func (c *Client) PushSingleAccount(deviceType int, account string, message Message) Response {
	params := make(map[string]interface{})

	params["account_list"] = account
	params["message"] = string(message.XGjson())
	res := c.push(deviceType,RESTAPI_PUSHSINGLEACCOUNT,params)
	return res
}
//推送给多个账户 最多100个
func (c *Client) PushAccountList(deviceType int, accountList []string, message Message) Response {
	params := make(map[string]interface{})

	account_list,err := json.Marshal(accountList)
	if(err != nil){

	}
	params["account_list"] = string(account_list)
	params["message"] = string(message.XGjson())
	res := c.push(deviceType,RESTAPI_PUSHACCOUNTLIST,params)
	return res
}
//推送给所有设备
func (c *Client) PushAllDevices(deviceType int, message Message) Response {
	params := make(map[string]interface{})

	params["message"] = string(message.XGjson())
	res := c.push(deviceType,RESTAPI_PUSHALLDEVICE,params)
	return res
}
//拼接公共参数
func (c *Client) push(deviceType int ,uri string, params map[string]interface{}) Response {
	switch deviceType {
	case Android:
		params["multi_pkg"] = 1 //0表示按注册时提供的包名分发消息；1表示按access id分发消息，所有以该access id成功注册推送的app均可收到消息。本字段对iOS平台无效
		params["message_type"] = 1 //消息类型：1：通知 2：透传消息。iOS平台请填0
	case Ios:
		params["message_type"] = 0 //消息类型：1：通知 2：透传消息。iOS平台请填0
		params["environment"] = 2	//向iOS设备推送时必填，1表示推送生产环境；2表示推送开发环境。推送Android平台不填或填0
	}
	params["expire_time"] = 300
	params["send_time"] = time.Now().Unix()
	params["timestamp"] = time.Now().Unix()
	params["access_id"] = c.accessId

	params["sign"] = c.generateSign(METHOD_POST, uri, c.secretKey, params)

	return c.send(uri, params)
}
//发送请求
func (c *Client) send(uri string, params map[string]interface{}) Response {
	var res Response = newResponse()
	data := make([]string, 0)
	for k, v := range params {
		data = append(data, fmt.Sprintf("%s=%v", k, v))
	}
	d := strings.Join(data, "&")
	fmt.Println(d)
	r, err := http.Post(uri, "application/x-www-form-urlencoded", strings.NewReader(d))
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	json.Unmarshal(body, &res)
	return res
}

func (c *Client) generateSign(method, uri, secretKey string, params map[string]interface{}) string {

	method = strings.ToUpper(method)
	u, err := url.Parse(uri)
	if err == nil {
		uri = fmt.Sprintf("%s%s", u.Host, u.Path)
	}

	param_str := make([]string, 0)
	keys := ksort(params)
	for _, k := range keys {
		param_str = append(param_str, fmt.Sprintf("%s=%v", k, params[k]))
	}

	origin := fmt.Sprintf("%s%s%s%s", method, uri, strings.Join(param_str, ""), secretKey)

	tmp := md5.Sum([]byte(origin))
	return hex.EncodeToString(tmp[:])
}

func NewClient(accessId int, secretKey string) *Client {
	return &Client{accessId, secretKey}
}

func ksort(p map[string]interface{}) []string {
	keys := make([]string, 0)
	for k, _ := range p {
		keys = append(keys, k)
	}
	list := sort.StringSlice(keys)
	sort.Sort(list)
	return []string(list)
}