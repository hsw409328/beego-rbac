package gofunc

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"gopkg.in/mgo.v2/bson"
	"io"
	"io/ioutil"
	mr "math/rand"
	"net"
	"net/http"
	"net/url"
	"os"
	"path"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// md5加密
func Md5Encrypt(str string) string {
	strMd5 := md5.New()
	strMd5.Write([]byte(str))
	return hex.EncodeToString(strMd5.Sum(nil))
}

// Sha1加密
func Sha1Encrypt(str string) string {
	strSha1 := sha1.New()
	strSha1.Write([]byte(str))
	return hex.EncodeToString(strSha1.Sum(nil))
}

//获取当日期
func CurrentDate() string {
	return time.Now().Format("2006-01-02")
}

//获取当前时间
func CurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 将日期格式化为时间戳
func StringToTime(strTime string) int64 {
	//获取本地location
	toBeCharge := strTime
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc)
	sr := theTime.Unix()

	return sr
}

// 将日期格式化为时间对象
func StringToTimeObject(strTime string, formatTpl string) time.Time {
	//获取本地location
	toBeCharge := strTime
	timeLayout := "2006-01-02 15:04:05"
	if formatTpl != "" {
		timeLayout = formatTpl
	}
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc)
	return theTime
}

// 将时间戳格式化为日期
func TimeUnixIntToString(intTime int64) string {
	timeLayout := "2006-01-02 15:04:05"
	dataTimeStr := time.Unix(intTime, 0).Format(timeLayout)
	return dataTimeStr
}

// 将时间戳格式化为自定义日期
func TimeUnixIntToStringCustom(intTime int64, formatTpl string) string {
	timeLayout := "2006-01-02 15:04:05"
	if formatTpl != "" {
		timeLayout = formatTpl
	}
	dataTimeStr := time.Unix(intTime, 0).Format(timeLayout) //设置时间戳 使用模板格式化为日期字符串
	return dataTimeStr
}

/**
* 获取以前的时间，时间戳
* @params sign string y,m,d 分别代表年，月，日
* @params num int 取多少天以前 或者 多少天以后 例如：-1 1天前 1 1天后
*
* @return unix.Time int64
*
 */
func DateFormatTime(needTime time.Time, sign string, num int) int64 {
	timeNow := needTime
	year, month, day := timeNow.Date()

	t := time.Date(year, month, day, timeNow.Hour(), timeNow.Minute(), timeNow.Second(), timeNow.Nanosecond(), timeNow.Location())
	var tmp time.Time
	switch sign {
	case "y":
		tmp = t.AddDate(num, 0, 0)
		break
	case "m":
		tmp = t.AddDate(0, num, 0)
		break
	case "d":
		tmp = t.AddDate(0, 0, num)
		break
	default:
		tmp = t.AddDate(0, 0, num)
		break
	}
	return tmp.Unix()
}

// uuid方法
func GetGuuid() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return Md5Encrypt(base64.URLEncoding.EncodeToString(b))
}

// 将未知类型转为字符串
func InterfaceToString(inter interface{}) (s string) {
	tempStr := ""
	switch inter.(type) {
	case nil:
		tempStr = ""
		break
	case string:
		tempStr = inter.(string)
		break
	case float64:
		tempStr = strconv.FormatFloat(inter.(float64), 'f', -1, 64)
		break
	case float32:
		tempStr = strconv.FormatFloat(float64(inter.(float32)), 'f', -1, 64)
		break
	case int64:
		tempStr = strconv.FormatInt(inter.(int64), 10)
		break
	case int:
		tempStr = strconv.Itoa(inter.(int))
		break
	case bool:
		tempStr = strconv.FormatBool(inter.(bool))
	case bson.ObjectId:
		tempStr = inter.(bson.ObjectId).Hex()
	case []interface{}:
		tempStr, _ = JsonToString(inter)
	case []int:
		tempStr, _ = JsonToString(inter)
	case []int64:
		tempStr, _ = JsonToString(inter)
	case []float32:
		tempStr, _ = JsonToString(inter)
	case []float64:
		tempStr, _ = JsonToString(inter)
	case map[string]interface{}:
		tempStr, _ = JsonToString(inter)
	case map[string]string:
		tempStr, _ = JsonToString(inter)
	case time.Time:
		tempStr = inter.(time.Time).String()
	default:
		tempStr = "Error! Not Found Type!"
	}
	return tempStr
}

func JsonToString(inter interface{}) (string, error) {
	by, err := json.Marshal(inter)
	if err != nil {
		return "", err
	} else {
		return string(by), nil
	}
}

// UTF8 Slice 转为 GBK Slice
func UTF8SliceToGBKSlice(strSlice []string) []string {
	strString := strings.Join(strSlice, ",")
	reader := transform.NewReader(bytes.NewReader([]byte(strString)), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return strSlice
	}
	return strings.Split(string(d), ",")
}

// UTF8 字符串 转为 GBK 字符串
func UTF8StringToGBKString(str string) string {
	reader := transform.NewReader(bytes.NewReader([]byte(str)), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return str
	}
	return string(d)
}

func ConvertToMap(model interface{}) bson.M {
	ret := bson.M{}

	modelReflect := reflect.ValueOf(model)

	if modelReflect.Kind() == reflect.Ptr {
		modelReflect = modelReflect.Elem()
	}

	modelRefType := modelReflect.Type()
	fieldsCount := modelReflect.NumField()

	var fieldData interface{}

	for i := 0; i < fieldsCount; i++ {
		field := modelReflect.Field(i)

		switch field.Kind() {
		case reflect.Struct:
			fallthrough
		case reflect.Ptr:
			fieldData = ConvertToMap(field.Interface())
		default:
			fieldData = field.Interface()
		}
		ret[modelRefType.Field(i).Name] = fieldData
	}

	return ret
}

/**
* 获取以前的时间，时间戳
* @params sign string y,m,d,h 分别代表年，月，日, 小时
* @params num int 取多少天以前 或者 多少天以后 例如：-1 1天前 1 1天后
*
* @return unix.Time int64
*
 */
func LastTime(sign string, num int) int64 {
	timeNow := time.Now()
	year, month, day := timeNow.Date()
	t := time.Date(year, month, day, 0, 0, 0, 0, timeNow.Location())
	var tmp time.Time
	switch sign {
	case "y":
		tmp = t.AddDate(num, 0, 0)
		break
	case "m":
		tmp = t.AddDate(0, num, 0)
		break
	case "d":
		tmp = t.AddDate(0, 0, num)
		break
	case "h":
		lastSecond := num * 60 * 60
		tmpUnixInt := timeNow.Unix() + int64(lastSecond)
		tmp = StringToTimeObject(TimeUnixIntToString(tmpUnixInt), "")
	default:
		tmp = t.AddDate(0, 0, num)
		break
	}
	return tmp.Unix()
}

// 获取本周的，周一到周日的时间
func GetWeekMondayAndSundayDateString() (string, string) {
	weekStr := time.Now().Weekday().String()
	preDay := 0
	nextDay := 0
	switch weekStr {
	case "Monday":
		preDay = 0
		nextDay = 6
		break
	case "Tuesday":
		preDay = -1
		nextDay = 5
		break
	case "Wednesday":
		preDay = -2
		nextDay = 4
		break
	case "Thursday":
		preDay = -3
		nextDay = 3
		break
	case "Friday":
		preDay = -4
		nextDay = 2
		break
	case "Saturday":
		preDay = -5
		nextDay = 1
		break
	case "Sunday":
		preDay = -6
		nextDay = 0
		break
	}
	nextDayTime := strings.Split(TimeUnixIntToString(LastTime("d", nextDay)), " ")
	return TimeUnixIntToString(LastTime("d", preDay)), nextDayTime[0] + " 23:59:59"
}

// 获取每年的1号到当前时间的每周开始与结束时间
func GetWeekStartAndEndTime() []interface{} {
	_, w := time.Now().ISOWeek()
	signalInt := 0
	var s, e string
	var listRes []interface{}
	for ; w > 0; w-- {
		if signalInt == 0 {
			//先获取本周的开始与结束时间
			s, e = GetWeekMondayAndSundayDateString()
			//填充当前周的时间
			listRes = append(listRes, map[string]string{"start": s, "end": e})
			signalInt = 1
		} else {
			s = TimeUnixIntToString(DateFormatTime(StringToTimeObject(s, ""), "d", -7))
			e = TimeUnixIntToString(DateFormatTime(StringToTimeObject(e, ""), "d", -7))
			//填充时间
			listRes = append(listRes, map[string]string{"start": s, "d": e})
		}
	}
	return listRes
}

// URI获取域名
func GetDomain(urlStr string) (string, error) {
	urlStr = strings.Replace(urlStr, " ", "", -1)
	urlStr = strings.Replace(urlStr, "　", "", -1)
	strRes, err := url.QueryUnescape(urlStr)
	if err != nil {
		return "", err
	}
	if !strings.Contains(urlStr, "http://") && !strings.Contains(urlStr, "https://") {
		strRes = "http://" + strRes
	}
	urlObject, err := url.Parse(strRes)
	if err != nil {
		return "", err
	}
	//判断是否为域名
	boolErr := IsDomain(urlObject.Host)
	if !boolErr {
		return "", errors.New("非域名")
	}
	return urlObject.Host, nil
}

// 将远程图片转换为base64字符串
func RemoteImageToBase64(imageUrl string) (string, error) {
	if len(imageUrl) == 0 {
		return "", errors.New("未找到图片地址")
	}
	extStr := strings.Replace(path.Ext(imageUrl), ".", "", -1)
	resp, err := http.Get(imageUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	by, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	base64Str := "data:image/" + extStr + ";base64," + base64.StdEncoding.EncodeToString(by)
	return base64Str, nil

}

// 获取两个日期相差多少天
func TimeSubDays(t1, t2 time.Time) int {
	if t1.Location().String() != t2.Location().String() {
		return -1
	}
	hours := t1.Sub(t2).Hours()

	if hours <= 0 {
		return -1
	}
	// sub hours less than 24
	if hours < 24 {
		// may same day
		t1y, t1m, t1d := t1.Date()
		t2y, t2m, t2d := t2.Date()
		isSameDay := (t1y == t2y && t1m == t2m && t1d == t2d)

		if isSameDay {
			return 0
		} else {
			return 1
		}
	} else { // equal or more than 24
		return int(hours / 24)
	}
}

// 判断obj是否在target中，target支持的类型arrary,slice,map
func ContainObjectInTarget(obj interface{}, target interface{}) (bool) {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}
	return false
}

// 判断字符串是否在另外一个字符串出现
func Strpos(str string, needStr interface{}) bool {
	needTmpStr := InterfaceToString(needStr)
	return strings.Contains(str, needTmpStr)
}

// 判断是否为域名
func IsDomain(str string) bool {
	p := regexp.MustCompile(`\A[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)+\z`)
	if p.MatchString(str) {
		return true
	} else {
		return false
	}
}

// 判断是否为IP
func IsIP(str string) bool {
	p := regexp.MustCompile(`\A(?:[0-9]{1,3}\.){3}[0-9]{1,3}\z`)
	if p.MatchString(str) {
		return true
	} else {
		return false
	}
}

// 判断是否为URL
func IsUrl(str string) bool {
	p := regexp.MustCompile(`\A(.+?)://[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)+(?::\d{1,5})?\S*\z`)
	if p.MatchString(str) {
		return true
	} else {
		return false
	}
}

// 判断是否为HOST
func IsHost(str string) bool {
	p := regexp.MustCompile(`\A[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)+(?::\d{1,5})\z`)
	if p.MatchString(str) {
		return true
	} else {
		return false
	}
}

// 字符类型转换成Map类型
func StringToMap(str string) (map[string]interface{}, error) {
	var jsonMap map[string]interface{}
	err := json.Unmarshal([]byte(str), &jsonMap)
	if err != nil {
		return nil, err
	}
	return jsonMap, nil
}

// 字符串类型转成Slice类型
func StringToSlice(str string) ([]interface{}, error) {
	var jsonMap []interface{}
	err := json.Unmarshal([]byte(str), &jsonMap)
	if err != nil {
		return nil, err
	}
	return jsonMap, nil
}

// Map Or Slice 装换为Json
func MapOrSliceToJsonString(mapData interface{}) (string, error) {
	byteBody, err := json.Marshal(mapData)
	if err != nil {
		return "", err
	}
	return string(byteBody), nil
}

// 获取本地IP
func GetLocalIp() string {
	ipStr := "127.0.0.1"
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ipStr
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipStr = ipnet.IP.String()
			}
		}
	}
	return ipStr
}

// 合并两个String-Slice
func SliceMerge(s1 []string, s2 []string) []string {
	slice := make([]string, len(s1)+len(s2))
	copy(slice, s1)
	copy(slice[len(s1):], s2)
	return slice
}

// 按行进行读取文件
func ReadLinesForFile(file string) ([]string, error) {
	var lines []string
	fi, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer fi.Close()
	br := bufio.NewReader(fi)
	for {
		line, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		lines = append(lines, string(line))
	}
	return lines, nil
}

// 获取随机数(8位)
func RandomString() string {
	RndInit := mr.New(mr.NewSource(time.Now().UnixNano()))
	rndStrResult := fmt.Sprintf("%08v", RndInit.Int31n(100000000))
	return rndStrResult
}

func GetCurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)

	return path.Dir(filename)
}

// 产生正则实体
func RegexpCompile(str string) *regexp.Regexp {
	return regexp.MustCompile("^" + str + "$")
}

// 判断val是否能正确匹配exp中的正则表达式。
// val可以是[]byte, []rune, string类型。
func RegexpIsMatch(exp *regexp.Regexp, val interface{}) bool {
	switch v := val.(type) {
	case []rune:
		return exp.MatchString(string(v))
	case []byte:
		return exp.Match(v)
	case string:
		return exp.MatchString(v)
	default:
		return false
	}
}

// 连接最后一个字符
func ConnectLastWord(oldString, lastWord string) string {
	if oldString == "" {
		return lastWord
	}
	oldSliceString := strings.Split(oldString, "")
	if oldSliceString[len(oldSliceString)-1] == lastWord {
		return oldString
	}
	return oldString + lastWord
}

// 连接第一个字符
func ConnectFirstWord(oldString, firstWord string) string {
	if oldString == "" {
		return firstWord
	}
	oldSliceString := strings.Split(oldString, "")
	if oldSliceString[0] == firstWord {
		return oldString
	}
	return firstWord + oldString
}

// 获取两个日期的区间值  支持秒级时间戳，毫秒级暂时不支持
func GetStartTimeAndLastTimeList(startTime, endTime int64) []string {
	var result = make([]string, 0)
	if endTime < startTime {
		return []string{time.Now().Format("2006-01-02 15:04:05")}
	}
	dayGap := int((endTime - startTime) / (24 * 60 * 60))
	//计算出相差天数
	//根据相差天数，从开始时间进行遍历获取
	//将startTime 转为time.Time类型
	t := time.Unix(startTime, 0)
	for i := 0; i <= dayGap; i++ {
		result = append(result, time.Unix(CustomLastTime(t, "d", i), 0).Format("2006-01-02"))
	}
	return result
}

/**
* 获取定制的的时间，时间戳
* @params customTime 自定义时间对象
* @params sign string y,m,d,h 分别代表年，月，日, 小时
* @params num int 取多少天以前 或者 多少天以后 例如：-1 1天前 1 1天后
*
* @return unix.Time int64
*
 */
func CustomLastTime(customTime time.Time, sign string, num int) int64 {
	timeNow := customTime
	year, month, day := timeNow.Date()
	t := time.Date(year, month, day, 0, 0, 0, 0, timeNow.Location())
	var tmp time.Time
	switch sign {
	case "y":
		tmp = t.AddDate(num, 0, 0)
		break
	case "m":
		tmp = t.AddDate(0, num, 0)
		break
	case "d":
		tmp = t.AddDate(0, 0, num)
		break
	case "h":
		lastSecond := num * 60 * 60
		tmpUnixInt := timeNow.Unix() + int64(lastSecond)
		tmp = StringToTimeObject(TimeUnixIntToString(tmpUnixInt), "")
	default:
		tmp = t.AddDate(0, 0, num)
		break
	}
	return tmp.Unix()
}
