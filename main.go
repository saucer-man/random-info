package randominfo

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Info struct {
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Addr   string `json:"addr"`
	IDcard string `json:"idcard"`
}

func RandomName() (string, error) {
	rand.Seed(time.Now().Unix())
	return fmt.Sprintf("%s%s", firstName[rand.Intn(len(firstName))], secondName[rand.Intn(len(secondName))]), nil
}

func RandomAddr() (string, error) {
	m := make(map[string]map[string]map[string][]string)
	err := json.Unmarshal([]byte(rawAddr), &m)
	if err != nil {
		return "", err
	}
	rand.Seed(time.Now().Unix())
	var province string
	var city string
	var area string
	var street string

	r := rand.Intn(len(m))
	for k := range m {
		if r == 0 {
			province = k
		}
		r--
	}

	r = rand.Intn(len(m[province]))
	for k := range m[province] {
		if r == 0 {
			city = k
		}
		r--
	}

	r = rand.Intn(len(m[province][city]))
	for k := range m[province][city] {
		if r == 0 {
			area = k
		}
		r--
	}
	streets := m[province][city][area]
	if len(streets) != 0 {
		street = streets[rand.Intn(len(streets))]
	}
	return fmt.Sprintf("%s %s %s %s", province, city, area, street), nil
}

func RandomPhone() (string, error) {
	rand.Seed(time.Now().Unix())
	phone := phonePrefix[rand.Intn(len(phonePrefix))] + fmt.Sprintf("%0*d", 8, rand.Intn(100000000))
	return phone, nil
}

func RandomIDcard(areacode string) (string, error) {
	rand.Seed(time.Now().Unix())
	if areacode == "" {
		r := rand.Intn(len(areaCode))
		for k := range areaCode {
			if r == 0 {
				areacode = k
			}
			r--
		}
	}
	// generate random Birthday
	t, err := RandomBirthday(true)
	if err != nil {
		return "", err
	}
	birthday := t.UTC().Format("20060102")
	randomCode := fmt.Sprintf("%0*d", 3, rand.Intn(999))
	prefix := areacode + birthday + randomCode
	code, err := VerifyCode(prefix)
	if err != nil {
		return "", err
	}
	idCard := prefix + code
	return idCard, err
}

// randBirthday isFullAge: true 年满18岁
func RandomBirthday(isFullAge bool) (time.Time, error) {
	var (
		begin, end time.Time
		err        error
	)
	if isFullAge {
		if begin, err = time.Parse("2006-01-02 15:04:05", time.Now().AddDate(-70, 0, 0).Format("2006-01-02 15:04:05")); err != nil {
			return time.Time{}, err
		}
		if end, err = time.Parse("2006-01-02 15:04:05", time.Now().AddDate(-18, 0, 0).Format("2006-01-02 15:04:05")); err != nil {
			return time.Time{}, err
		}

	} else {
		if begin, err = time.Parse("2006-01-02 15:04:05", "1970-01-01 00:00:00"); err != nil {
			return time.Time{}, err
		}
		if end, err = time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05")); err != nil {
			return time.Time{}, err
		}
	}
	rand.Seed(time.Now().UnixNano())
	return time.Unix(begin.UTC().Unix()+rand.Int63n(end.UTC().Unix()-begin.UTC().Unix()), 0), nil
}

// 获取 VerifyCode
func VerifyCode(cardId string) (ret string, err error) {
	var ValCodeArr = [...]string{
		"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2",
	}
	var Wi = [...]int{
		7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2,
	}
	tmp := 0
	for i, v := range Wi {
		if t, _err := strconv.Atoi(string(cardId[i])); _err == nil {
			tmp += t * v
		} else {
			err = _err
			return
		}
	}
	return ValCodeArr[tmp%11], nil
}

func RandomInfo() (info Info, err error) {
	info.Name, err = RandomName()
	if err != nil {
		return
	}
	info.Phone, err = RandomPhone()
	if err != nil {
		return
	}
	info.Addr, err = RandomAddr()
	if err != nil {
		return
	}
	// 根据地址生成身份证
	areaArray := strings.Split(info.Addr, " ")
	if len(areaArray) != 4 {
		err = fmt.Errorf("unexpected addr: %s", info.Addr)
		return
	}
	area := areaArray[0] + areaArray[1] + areaArray[2]
	areacode, ok := areaCode[area]
	if !ok {
		err = fmt.Errorf("addr No match areacode: %s", info.Addr)
		return
	}
	// info.Addr = strings.ReplaceAll(info.Addr, " ", "")
	info.IDcard, err = RandomIDcard(areacode)
	return
}
