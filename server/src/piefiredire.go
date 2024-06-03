package piefiredire

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

var httpGet = http.Get

func GetBeefSummary() map[string]int32 {
	// URL ที่ต้องการดึงข้อมูล
	url := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"

	// ทำการ request ไปยัง URL
	resp, err := httpGet(url)
	if err != nil {
		fmt.Printf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		// อ่านข้อมูลจาก response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Failed to read response body: %v", err)
		}

		// แปลงข้อมูลที่อ่านมาเป็นสตริง
		text := string(body)
		// text := "Fatback t-bone t-bone, pastrami  ..   t-bone.  pork, meatloaf jowl enim.  Bresaola t-bone."
		re := regexp.MustCompile(`[.,]`)
		text = re.ReplaceAllString(text, "")
		// // แยกคำโดยไม่สนใจช่องว่าง
		words := strings.Fields(text)
		wordMap := make(map[string]int32)
		// แสดงผลคำที่แยกได้
		for _, word := range words {
			word = strings.ToLower(word)
			wordMap[word] = wordMap[word] + 1
		}
		return wordMap
	}
	return nil
}
