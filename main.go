package randominfo

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

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
