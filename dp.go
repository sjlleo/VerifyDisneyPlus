package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func RequestIP(requrl string, ip string, method string) string {

	urlValue, err := url.Parse(requrl)
	if err != nil {
		return "400"
	}
	host := urlValue.Host
	if ip == "" {
		ip = host
	}
	newrequrl := strings.Replace(requrl, host, ip, 1)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{ServerName: host},
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
		Timeout:       5 * time.Second,
	}
	data := url.Values{"grant_type": {"refresh_token"}, "refresh_token": {"eyJ6aXAiOiJERUYiLCJraWQiOiJLcTYtNW1Ia3BxOXdzLUtsSUUyaGJHYkRIZFduRjU3UjZHY1h6aFlvZi04IiwiY3R5IjoiSldUIiwiZW5jIjoiQzIwUCIsImFsZyI6ImRpciJ9..OdwL8TEIFZouLDJe.wLz6zEC3PlPAGxx4X4qyP837lUbFrI_DQGnrJDMtEaQd5gsjHwaYshscoDXCYjMioU8JvsH_HKZga3fzSDEoWuMA5lgv4dyJpoB4Cqi91JjPSkqsRHKZ1I-nRoTmnSkcW3RHE-0coAqDWgK7IZ5cPiHQ-9KVRqqZkmTbEHynBdgH2y-FJP8zK0-dAynzR2krlUahhcykp7J7VqhZj_l5HVZZkPylZ6eKoK4J8fQvuGJoqMaRZTzrIH4Yk9J3GMbKnYqEG3SKRp5qAuWTtqLDOoGN0wWsUE5VRuCZxRKpxayJWABq2u4ABkAtIqUx8CPx77ZXxZVlcjRN1Xa8F2-e2mTxZq_1FgzmWECFg6onkDj_TpfBdeFoxDzhnRNceoQ-iyyNf3sgxJ_nz_bwztVZf0Vt3OR8yBnXfbkuEY7GQ4pvCuy-peW0mwJJCd2eJ9ADwDEGmoY4F47W-8rxdBhgna-0hu0FuLxt9MlmH_tGCmM_T-61xsxymLO9tlkwBnxNw4u6T9X2hcvC7-4uzr5cJiaJ3sGPMNo_ixTrP8SG9zCIse-X6_Lq0v3Uo-QOKhcD4N3gIfwZFYEvf-HVGWzFpU683q9CJfTTEXhsufj1URhSis7GdAa3nLZVt7CScsMPcYrMI317PmU-Brdvl_Ic4QeHTeF8-57kzD3mm5mrlQ7kQIXQzzQPqHYt70MzxL_scfT90cpYaSOBQnB1l--226h7X51XxSbrOcO-25zS7OSyedya8eMG6zAmgkk1zvZUzdCHZyzYD8-t0KYcfA5AwiLIFHxgqL4ni9fVy-SpYTKRwCmkp_pZOPaFwJh8zkhw8QaSLHq7ubko7H1kjJZxzsG1l4Bla1QRlj_-FVoY8GZ6okFk3Ts6A2qOK6v8UT7sL_w2zaHDQH1q2o05vsLwqIOxg3Xyey0tahzPbl-In_i1JGGvqGXOiPcKL5uOcTOo1luk32AbCS9i5mkopTS401YYYMH-Sx_krW_VJd2czpFefc0dlagtzBytqlcyscscFwq6IE6VHwG2Ij-WfO44G5hGDJFkZMZLeDUnTIyNrLe9hcfJp73koOSFnURsFWFjM2lgUIayiREAl02oh2alUyqnG09gdXufT_2W0DjA4i7qYuv6ol5NIVc389dF3x4a_7dPBvsMU3ppA1rlV04FlK6_fRv-Dk_jclXRZiQ5ul2ZO2CQ96LmrzmkdeNxFxcwaNXCJGBiRWXfMunoddIRg_LrVGuqWRgxj4DEnngZ2-qI_dliGiYraIehsHvtWeXIUWNF_FQSnQgZLg4WPekcluCecE4Iv7Sz36k9GUDqqs8hRWddirhufYem6RC84PyNqafCnwczrx5pOacmVzDl9Oi8OIhdDasdJa7gvsDoFzf6bv5st7EvbORkgPs6MK46mDMlwkL7TqjrJnSJzozCX4zLbYeyiWK6EXCehOpImMN262KLYQxnf5ugvk11gIA4NXpTbzyo4hp2LS7u8UMs5_w3t02vizxSQGokp-3qkEWmViy3pup1IXMPrcpS6KWHX0AYi1oRDZB5B8vM04pRHwYjsgMp2L-w4PMaDC4QDRU81IdvQ5VRkyLT7CL5hDlq5smXw_7wSFTWxs9vc5PmnrykSAkwFPocORC2j4T96uiu3z4gNoBu_dwKNcPi-dV7myC4iRRTmpm0V5A9IW510RGTyso_b-1hUeGvToYl9VwNgN7Impt3PjEQO2HXMU3p96tdulDEA_8bbyPdEGfxxVK3k2n_dxj_GzPKA8V4ESoNMRrV1vCuxPnrzfAOhqmNOEewTHqlxENSsZFGvfzVj1KemR7zLky14JMVslILnvxl6vuX7SbfIQ5JDktq9qKtTKo1mFrA-mBS3n00FacjPi364nnugiWQN7EwhNdEDH_KtWXGZVh-u2NM5cdoS1kAsOKSLxFTnTDG738LhoB3i_ZOjHFASKiZcsX6yD5csIP21jG5nFF9Qw2qsnqmxRuDLilIoGczEMt2Pfo180CG8Dyr7XtOYNeVU7__h9zBm9CvaAHDoQQhU4KlXM4LsljFeajw5f2wn08OmsdfkSYYl45O718QgzR_RRqwDpQH2pyKDJZ9yZt5OCyxcbnCgepjUyp6S-Pigfw73ASoCknhLLheb2mqkWIC-s3NmClpMoK-IyE57AiHHCatZfPGPnNofVioN5SbVR08mV7pdyQEhQGxGFM_LTAFFpwC48gOFTq-FWdV58muDULTqO3ImbGG6X3vV-PVbher1oJx0CFnelGGIx9lwM-yHbpVZGq9IXnKqoblCHiwuaJgbCKBnTjia2gYPNlN0Ql1ia3vQc7bybDVHyLePAVbOk10MdwHprwMGE__wsXqagElQCGJpU3ytPDktncRPCSQBQ3mw94CCIOQYEyhnA1Vik127AznwbR10Xm59diGBtix0Ao-VIrjKzQNw2hXqC_H-IgY46OT5ZndZ02SAe6AVyipq6kTui_ZyuQhy-zAOiat4t6qh-LyL1xImBuOZ7e79737LYiLHEIgHOIQ68DKcSmsIuA.gwrRhM5AiYUQ6iAbRZhxlw"}, "subject_token_type": {"urn:bamtech:params:oauth:token-type:device"}}
	req, err := http.NewRequest("POST", newrequrl, strings.NewReader(data.Encode()))
	if err != nil {
		return "400"
	}
	req.Host = host
	req.Header.Set("USER-AGENT", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36")
	req.Header.Set("authorization", "Bearer ZGlzbmV5JmJyb3dzZXImMS4wLjA.Cu56AgSfBTDag5NiRA81oLHkDZfu5L3CKadnefEAY84")
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return "400"
	}
	defer resp.Body.Close()
	switch method {
	case "auth":
		s, _ := ioutil.ReadAll(resp.Body)
		response := string(s)

		if strings.Index(response, "unauthorized") != -1 {
			return "unauthorized"
		}
		return "ok"
	case "query":
		Header := resp.Header
		//fmt.Print(Header)
		if Header["Location"] == nil {
			StartIndex := strings.Index(Header["Set-Cookie"][0], "y=")
			EndIndex := strings.Index(Header["Set-Cookie"][0], ";")
			return Header["Set-Cookie"][0][StartIndex+2 : EndIndex]
		} else if Header["Location"][0] == "https://disneyplus.disney.co.jp/" {
			return "JP"
		} else if Header["Location"][0] == "https://preview.disneyplus.com/unavailable/" {
			return "Unavailable"
		}
	}
	return "-1"
}

func VerifyAuthorized() int {
	ipv4, _ := lookupIP("global.edge.bamgrid.com")

	tokenStatusv4 := RequestIP("https://global.edge.bamgrid.com/token", ipv4, "auth")
	if tokenStatusv4 == "ok" {
		return 1
	} else if tokenStatusv4 == "400" {
		return -2
	} else {
		return -1
	}
}

func QueryAreaAvailable(protocol string) string {
	ipv4, ipv6 := lookupIP("www.disneyplus.com")
	switch protocol {
	case "ipv4":
		return FindCountry(RequestIP("https://www.disneyplus.com", ipv4, "query"))
	case "ipv6":
		return FindCountry(RequestIP("https://www.disneyplus.com", ipv6, "query"))
	}
	return ""
}

func FindCountry(Code string) string {
	countryName := []string{"无信息", "美国", "阿富汗", "奥兰群岛", "阿尔巴尼亚", "阿尔及利亚", "美属萨摩亚", "安道尔", "安哥拉", "安圭拉", "南极洲", "安提瓜和巴布达", "阿根廷", "亚美尼亚", "阿鲁巴", "澳大利亚", "奥地利", "阿塞拜疆", "巴哈马", "巴林", "孟加拉国", "巴巴多斯", "白俄罗斯", "比利时", "伯利兹", "贝宁", "百慕大", "不丹", "玻利维亚", "波黑", "博茨瓦纳", "布维岛", "巴西", "英属印度洋领地", "文莱", "保加利亚", "布基纳法索", "布隆迪", "柬埔寨", "喀麦隆", "加拿大", "佛得角", "开曼群岛", "中非", "乍得", "智利", "中国", "圣诞岛", "科科斯（基林）群岛", "哥伦比亚", "科摩罗", "刚果（布）", "刚果（金）", "库克群岛", "哥斯达黎加", "科特迪瓦", "克罗地亚", "古巴", "塞浦路斯", "捷克", "丹麦", "吉布提", "多米尼克", "多米尼加", "厄瓜多尔", "埃及", "萨尔瓦多", "赤道几内亚", "厄立特里亚", "爱沙尼亚", "埃塞俄比亚", "福克兰群岛（马尔维纳斯）", "法罗群岛", "斐济", "芬兰", "法国", "法属圭亚那", "法属波利尼西亚", "法属南部领地", "加蓬", "冈比亚", "格鲁吉亚", "德国", "加纳", "直布罗陀", "希腊", "格陵兰", "格林纳达", "瓜德罗普", "关岛", "危地马拉", "格恩西岛", "几内亚", "几内亚比绍", "圭亚那", "海地", "赫德岛和麦克唐纳岛", "梵蒂冈", "洪都拉斯", "香港", "匈牙利", "冰岛", "印度", "印度尼西亚", "伊朗", "伊拉克", "爱尔兰", "英国属地曼岛", "以色列", "意大利", "牙买加", "日本", "泽西岛", "约旦", "哈萨克斯坦", "肯尼亚", "基里巴斯", "朝鲜", "韩国", "科威特", "吉尔吉斯斯坦", "老挝", "拉脱维亚", "黎巴嫩", "莱索托", "利比里亚", "利比亚", "列支敦士登", "立陶宛", "卢森堡", "澳门", "前南马其顿", "马达加斯加", "马拉维", "马来西亚", "马尔代夫", "马里", "马耳他", "马绍尔群岛", "马提尼克", "毛利塔尼亚", "毛里求斯", "马约特", "墨西哥", "密克罗尼西亚联邦", "摩尔多瓦", "摩纳哥", "蒙古", "黑山", "蒙特塞拉特", "摩洛哥", "莫桑比克", "缅甸", "纳米比亚", "瑙鲁", "尼泊尔", "荷兰", "荷属安的列斯", "新喀里多尼亚", "新西兰", "尼加拉瓜", "尼日尔", "尼日利亚", "纽埃", "诺福克岛", "北马里亚纳", "挪威", "阿曼", "巴基斯坦", "帕劳", "巴勒斯坦", "巴拿马", "巴布亚新几内亚", "巴拉圭", "秘鲁", "菲律宾", "皮特凯恩", "波兰", "葡萄牙", "波多黎各", "卡塔尔", "留尼汪", "罗马尼亚", "俄罗斯联邦", "卢旺达", "圣赫勒拿", "圣基茨和尼维斯", "圣卢西亚", "圣皮埃尔和密克隆", "圣文森特和格林纳丁斯", "萨摩亚", "圣马力诺", "圣多美和普林西比", "沙特阿拉伯", "塞内加尔", "塞尔维亚", "塞舌尔", "塞拉利昂", "新加坡", "斯洛伐克", "斯洛文尼亚", "所罗门群岛", "索马里", "南非", "南乔治亚岛和南桑德韦奇岛", "西班牙", "斯里兰卡", "苏丹", "苏里南", "斯瓦尔巴岛和扬马延岛", "斯威士兰", "瑞典", "瑞士", "叙利亚", "台湾", "塔吉克斯坦", "坦桑尼亚", "泰国", "东帝汶", "多哥", "托克劳", "汤加", "特立尼达和多巴哥", "突尼斯", "土耳其", "土库曼斯坦", "特克斯和凯科斯群岛", "图瓦卢", "乌干达", "乌克兰", "阿联酋", "英国", "美国本土外小岛屿", "乌拉圭", "乌兹别克斯坦", "瓦努阿图", "委内瑞拉", "越南", "英属维尔京群岛", "美属维尔京群岛", "瓦利斯和富图纳", "西撒哈拉", "也门", "赞比亚", "津巴布韦"}
	countryCode := []string{"null", "us", "af", "ax", "al", "dz", "as", "ad", "ao", "ai", "aq", "ag", "ar", "am", "aw", "au", "at", "az", "bs", "bh", "bd", "bb", "by", "be", "bz", "bj", "bm", "bt", "bo", "ba", "bw", "bv", "br", "io", "bn", "bg", "bf", "bi", "kh", "cm", "ca", "cv", "ky", "cf", "td", "cl", "cn", "cx", "cc", "co", "km", "cg", "cd", "ck", "cr", "ci", "hr", "cu", "cy", "cz", "dk", "dj", "dm", "do", "ec", "eg", "sv", "gq", "er", "ee", "et", "fk", "fo", "fj", "fi", "fr", "gf", "pf", "tf", "ga", "gm", "ge", "de", "gh", "gi", "gr", "gl", "gd", "gp", "gu", "gt", "gg", "gn", "gw", "gy", "ht", "hm", "va", "hn", "hk", "hu", "is", "in", "id", "ir", "iq", "ie", "im", "il", "it", "jm", "jp", "je", "jo", "kz", "ke", "ki", "kp", "kr", "kw", "kg", "la", "lv", "lb", "ls", "lr", "ly", "li", "lt", "lu", "mo", "mk", "mg", "mw", "my", "mv", "ml", "mt", "mh", "mq", "mr", "mu", "yt", "mx", "fm", "md", "mc", "mn", "me", "ms", "ma", "mz", "mm", "na", "nr", "np", "nl", "an", "nc", "nz", "ni", "ne", "ng", "nu", "nf", "mp", "no", "om", "pk", "pw", "ps", "pa", "pg", "py", "pe", "ph", "pn", "pl", "pt", "pr", "qa", "re", "ro", "ru", "rw", "sh", "kn", "lc", "pm", "vc", "ws", "sm", "st", "sa", "sn", "rs", "sc", "sl", "sg", "sk", "si", "sb", "so", "za", "gs", "es", "lk", "sd", "sr", "sj", "sz", "se", "ch", "sy", "tw", "tj", "tz", "th", "tl", "tg", "tk", "to", "tt", "tn", "tr", "tm", "tc", "tv", "ug", "ua", "ae", "gb", "um", "uy", "uz", "vu", "ve", "vn", "vg", "vi", "wf", "eh", "ye", "zm", "zw"}
	for i, v := range countryCode {
		if strings.Contains(Code, strings.ToUpper(v)) {
			return countryName[i]
		}
	}
	return Code
}

func lookupIP(domain string) (string, string) {
	var ipv4, ipv6 string
	ns, err := net.LookupHost(domain)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s", err.Error())
		return "Error", "Error"
	}

	switch {
	case len(ns) != 0:
		for _, n := range ns {

			if ParseIP(n) == 4 {
				ipv4 = n
			}
			if ParseIP(n) == 6 {
				ipv6 = "[" + n + "]"
			}
		}
	}
	return ipv4, ipv6
}

func ParseIP(s string) int {
	ip := net.ParseIP(s)
	if ip == nil {
		return 0
	}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '.':
			return 4
		case ':':
			return 6
		}
	}
	return 0
}

func main() {

	var NextLineSignal bool = false
	QueryStatusv4 := QueryAreaAvailable("ipv4")
	QueryStatusv6 := QueryAreaAvailable("ipv6")

	fmt.Println("** DisneyPlus 检测小工具 v1.0 Beta By \033[1;36m@sjlleo\033[0m **")

	VerifyStatus := VerifyAuthorized()
	if VerifyStatus == -2 {
		fmt.Println("\033[0;35m[提醒] 无法获取DisneyPlus权验接口信息，当前测试可能会不准确\033[0m")
	}

	switch QueryStatusv4 {
	case "400":
		break
	case "Unavailable":
		NextLineSignal = true
		fmt.Println("\033[0;36m[IPv4]\033[0m\n\033[0;33m当前IPv4出口不在DisneyPlus所支持的地区\033[0m")
		break
	case "-1":
		NextLineSignal = true
		fmt.Println("\033[0;36m[IPv4]\033[0m\n\033[0;34m当前IPv4出口所在地区即将开通DisneyPlus，尽请期待哦！\033[0m")
		break
	default:
		NextLineSignal = true
		fmt.Println("\033[0;36m[IPv4]\033[0m")
		if VerifyStatus == -1 {
			fmt.Println("\033[0;33m当前IPv4出口不能解锁DisneyPlus\033[0m")
		} else {
			fmt.Println("\033[0;32m当前IPv4出口解锁DisneyPlus\033[0m\n区域：\033[1;36m" + QueryStatusv4 + "区\033[0m")
		}
	}

	switch QueryStatusv6 {
	case "400":
		break
	case "Unavailable":
		if NextLineSignal == true {
			fmt.Print("\n")
		}
		fmt.Println("\033[0;36m[IPv4]\033[0m\n\033[0;33m当前IPv6出口不在DisneyPlus所支持的地区\033[0m")
		break
	case "-1":
		if NextLineSignal == true {
			fmt.Print("\n")
		}
		fmt.Println("\033[0;36m[IPv6]\033[0m\n\033[0;34m当前IPv6出口所在地区即将开通DisneyPlus，尽请期待哦！\033[0m")
		break
	default:
		if NextLineSignal == true {
			fmt.Print("\n")
		}
		fmt.Println("\033[0;36m[IPv6]\033[0m")
		if VerifyStatus == -1 {
			fmt.Println("\033[0;33m当前IPv6出口不能解锁DisneyPlus\033[0m")
		} else {
			fmt.Println("\033[0;32m当前IPv6出口解锁DisneyPlus\033[0m\n区域：\033[1;36m" + QueryStatusv6 + "区\033[0m")
		}
	}
	return
}
