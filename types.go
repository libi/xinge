package xinge

import (
	"encoding/json"
	"strings"
)

const (
	ACTION_TYPE_ACTIVITY int = 1
	ACTION_TYPE_URL      int = 2
	ACTION_TYPE_INTENT   int = 3

	Android int = 1
	Ios int = 2

	MESSAGE_TYPE_NOTIFICATION int = 1
	MESSAGE_TYPE_MESSAGE      int = 2

	METHOD_POST string = "post"

	RESTAPI_PUSHSINGLEDEVICE        string = "http://openapi.xg.qq.com/v2/push/single_device"
	RESTAPI_PUSHSINGLEACCOUNT       string = "http://openapi.xg.qq.com/v2/push/single_account"
	RESTAPI_PUSHACCOUNTLIST         string = "http://openapi.xg.qq.com/v2/push/account_list"
	RESTAPI_PUSHALLDEVICE           string = "http://openapi.xg.qq.com/v2/push/all_device"
	RESTAPI_PUSHTAGS                string = "http://openapi.xg.qq.com/v2/push/tags_device"
	RESTAPI_QUERYPUSHSTATUS         string = "http://openapi.xg.qq.com/v2/push/get_msg_status"
	RESTAPI_QUERYDEVICECOUNT        string = "http://openapi.xg.qq.com/v2/application/get_app_device_num"
	RESTAPI_QUERYTAGS               string = "http://openapi.xg.qq.com/v2/tags/query_app_tags"
	RESTAPI_CANCELTIMINGPUSH        string = "http://openapi.xg.qq.com/v2/push/cancel_timing_task"
	RESTAPI_BATCHSETTAG             string = "http://openapi.xg.qq.com/v2/tags/batch_set"
	RESTAPI_BATCHDELTAG             string = "http://openapi.xg.qq.com/v2/tags/batch_del"
	RESTAPI_QUERYTOKENTAGS          string = "http://openapi.xg.qq.com/v2/tags/query_token_tags"
	RESTAPI_QUERYTAGTOKENNUM        string = "http://openapi.xg.qq.com/v2/tags/query_tag_token_num"
	RESTAPI_CREATEMULTIPUSH         string = "http://openapi.xg.qq.com/v2/push/create_multipush"
	RESTAPI_PUSHACCOUNTLISTMULTIPLE string = "http://openapi.xg.qq.com/v2/push/account_list_multiple"
	RESTAPI_PUSHDEVICELISTMULTIPLE  string = "http://openapi.xg.qq.com/v2/push/device_list_multiple"
	RESTAPI_QUERYINFOOFTOKEN        string = "http://openapi.xg.qq.com/v2/application/get_app_token_info"
	RESTAPI_QUERYTOKENSOFACCOUNT    string = "http://openapi.xg.qq.com/v2/application/get_app_account_tokens"
)

//安卓推送结构定义
type TimeInterval struct {
	StartHour int `json:"startHour"` //range [0, 23]
	StartMin  int `json:"startMin"`  //range [0, 59]
	EndHour   int `json:"endHour"`   //range [0, 23]
	EndMin    int `json:"endMin"`    //range [0, 59]
}

type ClickAction struct {
	ActionType               int    `json:"actionType"`
	Url                      string `json:"url"`
	ConfirmOnUrl             int    `json:"confirmOnUrl"` // 1:yes, 0:no [default 0]
	Activity                 string `json:"activity"`
	AtyAttrIntentFlag        int    `json:"atyAttrIntentFlag"`
	AtyAttrPendingIntentFlag int    `json:"atyAttrPendingIntentFlag"`
	Intent                   string `json:"intent"`
}

type Style struct {
	BuilderId int    `json:"builderId"`
	Ring      int    `json:"ring"`
	Vibrate   int    `json:"vibrate"`
	Clearable int    `json:"clearable"`
	NId       int    `json:"nId"`
	Lights    int    `json:"lights"`
	IconType  int    `json:"iconType"`
	IconRes   string `json:"iconRes"`
	RingRaw   string `json:"ringRaw"`
	StyleId   int    `json:"styleId"`
	SmallIcon string `json:"smallIcon"`
}

type Message struct {
	Title        string            `json:"title"`
	Content      string            `json:"content"`
	AcceptTime   []TimeInterval    `json:"acceptTime"`
	Type         int               `json:"type"`
	Style        Style             `json:"style"`
	Action       ClickAction       `json:"action"`
	Custom       map[string]string `json:"custom"`
	LoopInterval int               `json:"loopInterval"` //unit:day, range:[1, 14]
	LoopTimes    int               `json:"loopTimes"`    // range:[1, 15]
}

type Response struct {
	Code int    `json:"ret_code"`
	Msg  string `json:"err_msg"`
}

func newResponse() Response {
	return Response{-1, "message not valid"}
}

type TagTokenPair struct {
	Tag   string `json:"tag"`
	Token string `json:"token"`
}

func NewMessage() *Message {
	return &Message{
		AcceptTime: make([]TimeInterval, 0),
		Style:      Style{BuilderId: 0},
		Action:     ClickAction{},
	}
}

func (m *Message) SetAction(action ClickAction) {
	m.Action = action
}

func (m *Message) SetStyle(style Style) {
	m.Style = style
}

func (m *Message) SetCustom(custom map[string]string) {
	m.Custom = custom
}

func (m *Message) AddAcceptTime(acceptTime TimeInterval) {
	m.AcceptTime = append(m.AcceptTime, acceptTime)
}

func (m *Message) XGjson() []byte {
	result := make(map[string]interface{})

	result["title"] = m.Title
	result["content"] = m.Content
	result["accept_time"] = m.JsonAcceptTime()

	if m.Type == MESSAGE_TYPE_NOTIFICATION {
		result["builder_id"] = m.Style.BuilderId
		result["ring"] = m.Style.Ring
		result["vibrate"] = m.Style.Vibrate
		result["clearable"] = m.Style.Clearable
		result["n_id"] = m.Style.NId

		if !strings.EqualFold(m.Style.RingRaw, "") {
			result["ring_raw"] = m.Style.RingRaw
		}

		result["lights"] = m.Style.Lights
		result["icon_type"] = m.Style.IconType

		if !strings.EqualFold(m.Style.IconRes, "") {
			result["icon_res"] = m.Style.IconRes
		}

		result["style_id"] = m.Style.StyleId

		if !strings.EqualFold(m.Style.SmallIcon, "") {
			result["small_icon"] = m.Style.SmallIcon
		}

		result["action"] = m.JsonAction()

	}

	result["custom_cnotent"] = m.Custom

	ret, err := json.Marshal(result)
	if err != nil {
		return nil
	}

	return ret
}

func (m *Message) JsonAcceptTime() []map[string]map[string]int {
	result := make([]map[string]map[string]int, 0)
	for _, t := range m.AcceptTime {
		var tmp map[string]map[string]int = map[string]map[string]int{
			"start": {"hour": t.StartHour, "min": t.StartMin},
			"end":   {"hour": t.EndHour, "min": t.EndHour},
		}
		result = append(result, tmp)
	}
	return result
}

func (m *Message) JsonAction() map[string]interface{} {
	result := make(map[string]interface{})
	result["action_type"] = m.Action.ActionType
	result["browser"] = map[string]interface{}{"url": m.Action.Url, "confirm": m.Action.ConfirmOnUrl}
	result["activity"] = m.Action.Activity
	result["intent"] = m.Action.Intent

	aty_attr := make(map[string]int)
	if m.Action.AtyAttrIntentFlag != 0 {
		aty_attr["if"] = m.Action.AtyAttrIntentFlag
	}

	if m.Action.AtyAttrPendingIntentFlag != 0 {
		aty_attr["pf"] = m.Action.AtyAttrPendingIntentFlag
	}

	result["aty_attr"] = aty_attr
	return result
}

//ios推送定义

type Aps struct {
	Alert string
	Badge int
}

type IosMessage struct {
	Aps Aps `json:"aps"`
	AcceptTime   []TimeInterval    `json:"acceptTime"`
	Custom       map[string]string `json:"custom"`
}

func NewIosMessage() *IosMessage {
	return &IosMessage{
		AcceptTime: make([]TimeInterval, 0),
		Custom : make(map[string]string, 0),
		Aps : Aps{},
	}
}
func (i *IosMessage)SetAps(title string,badge int){
	aps := Aps{Alert:title,Badge:badge}
	i.Aps = aps
}
func (i *IosMessage) AddAcceptTime(acceptTime TimeInterval) {
	i.AcceptTime = append(i.AcceptTime, acceptTime)
}

func (i *IosMessage)XGjson() []byte {
	result := make(map[string]interface{})

	result["aps"] = i.Aps

	result["custom1"] = i.Custom

	ret, err := json.Marshal(result)
	if err != nil {
		return nil
	}

	return ret
}

type CommonMessage interface {
	XGjson()[]byte
}

func DefaultMessage(title ,content string) *Message{
	message := NewMessage()
	message.Title = title
	message.Content = content

	message.Type = MESSAGE_TYPE_NOTIFICATION

	style := Style{BuilderId: 0, Ring: 1, Vibrate: 1, Clearable: 0, NId: 0}
	action := ClickAction{}
	action.ActionType = ACTION_TYPE_ACTIVITY
	message.SetStyle(style)
	message.SetAction(action)
	message.AddAcceptTime(TimeInterval{0, 0, 23, 59})
	return message
}

func DefaultIosMessage(title string,badge int) *IosMessage{
	iosmessage := NewIosMessage()
	iosmessage.Aps.Alert = title
	iosmessage.Aps.Badge = badge
	iosmessage.AddAcceptTime(TimeInterval{0, 0, 23, 59})
	return iosmessage
}


