package xinge

func PushSingleAndroidDevice(accessId int, secretKey, deviceToken, title, content string,custom map[string]string) Response {
	client := NewClient(accessId, secretKey)
	message := DefaultMessage(title,content)

	message.SetCustom(custom)
	res := client.PushSingleDevice(1,deviceToken, message)
	return res

}


func PushSingleIosDevice(accessId int, secretKey, deviceToken, title string,custom map[string]string) Response {
	client := NewClient(accessId, secretKey)
	message := DefaultIosMessage(title,1)
	message.Custom = custom
	res := client.PushSingleDevice(2,deviceToken, message)
	return res
}
