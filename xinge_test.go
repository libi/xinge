package xinge

import (
	"fmt"
	"testing"
)

var (
	accessId    int    = 2100146994
	secretKey   string = "d7e0de4cc42f6a33b84d0beaeabee1fe"
	deviceToken string = "e5e665ed947c8ba14a8ea78fa5b9b7dbc5ffed2e"
	account string = "uid123"
	account_list []string = []string{"uid2","uid123","uid21"}
	client *Client = NewClient(accessId,secretKey)
)

func TestStaticPushSingleADevice(t *testing.T) {
	res := client.PushSingleAndroidDevice(deviceToken, "有人偷了你的菜。", "",nil)
	if res.Code != 0 {
		t.Errorf("send failure, error is %s", res.Msg)
	}else{
		fmt.Println(res)
		fmt.Println("1 success")
	}
}
func TestStaticPushSingleIosDevice(t *testing.T) {
	res := client.PushSingleIosDevice(deviceToken, "有人偷了你的菜。",2,nil)
	if res.Code != 0 {
		t.Errorf("send failure, error is %s", res.Msg)
	}else{
		fmt.Println(res)
		fmt.Println("2 success")
	}
}

func TestStaticPushSingleAaccount(t *testing.T){
	res := client.PushSingleAndroidAccount(account, "有人偷了你的菜。","",nil)
	if res.Code != 0 {
		t.Errorf("send failure, error is %s", res.Msg)
	}else{
		fmt.Println(res)
		fmt.Println("3 success")
	}
}
func TestStaticPushSingleIosaccount(t *testing.T){
	res := client.PushSingleIosAccount(account, "有人偷了你的菜。",1,nil)
	if res.Code != 0 {
		t.Errorf("send failure, error is %s", res.Msg)
	}else{
		fmt.Println(res)
		fmt.Println("4 success")
	}
}

func TestStaticPushAAL(t *testing.T){
	res := client.PushAndroidAccountL(account_list, "有人偷了你的菜。","",nil)
	if res.Code != 0 {
		t.Errorf("send failure, error is %s", res.Msg)
	}else{
		fmt.Println(res)
		fmt.Println("5 success")
	}
}

func TestStaticPushIosAL(t *testing.T){
	res := client.PushIosAccountL(account_list, "有人偷了你的菜。",1,nil)
	if res.Code != 0 {
		t.Errorf("send failure, error is %s", res.Msg)
	}else{
		fmt.Println(res)
		fmt.Println("6 success")
	}
}
