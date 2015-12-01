package xinge

func (c *Client)PushSingleAndroidDevice(deviceToken, title, content string,custom map[string]string) Response {

	message := DefaultMessage(title,content)

	message.SetCustom(custom)
	return c.PushSingleDevice(Android,deviceToken, message)

}

func (c *Client)PushSingleIosDevice(deviceToken ,title string,badge int,custom map[string]string) Response {

	message := DefaultIosMessage(title,badge)
	message.Custom = custom
	return c.PushSingleDevice(Ios,deviceToken, message)

}

func (c *Client)PushSingleAndroidAccount(account,title,content string,custom map[string]string) Response{
	message := DefaultMessage(title,content)
	message.SetCustom(custom)
	return c.PushSingleAccount(Android,account,message)
}
func (c *Client)PushSingleIosAccount(account,title string,badge int,custom map[string]string) Response{
	message := DefaultIosMessage(title,badge)
	message.Custom = custom
	return c.PushSingleAccount(Ios,account,message)
}
func (c *Client)PushAndroidAccountL(accountlist []string,title,content string,custom map[string]string) Response{
	message := DefaultMessage(title,content)
	message.SetCustom(custom)
	return c.PushAccountList(Android,accountlist,message)
}
func (c *Client)PushIosAccountL(accountlist []string,title string,badge int,custom map[string]string) Response{
	message := DefaultIosMessage(title,badge)
	message.Custom = custom
	return c.PushAccountList(Ios,accountlist,message)
}

//todo