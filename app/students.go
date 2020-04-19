package app

import (
	"RedRock-web-back-end-2020-5-lv2/database"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

const (
	HEADEINFO  = "<li>〉〉(.*?学年)(\\d学期) 学生课表>>(\\d{10})(.*?)  <\\/li>"
	CLASSTABLE = "<div id=\"kbStuTabs-list\">(.*?|\\s*)*<\\/div>"
	CLASSINFO  = "<td ?>(.*?)(<\\/td>|>)\\s*<td>(星期\\d) ?(第.*?节) ?(.*?周) ?<\\/td><td>(.*?)(<\\/td>|>)"
)

var ch chan int

func GetAllStudentsInfo() {
	for i := 2019211100; i < 2019215111; i++ {
		go GetAStudentInfo(i)
	}
	<-ch
}

func GetAStudentInfo(id int) {
	var student database.Student
	var class database.Class

	student.StudentId = id
	class.StudentId = id

	url := "http://jwc.cqupt.edu.cn/kebiao/kb_stu.php?xh=" + strconv.Itoa(student.StudentId)
	body := GetBody(url)

	headInfoReg := regexp.MustCompile(HEADEINFO)
	headInfo := headInfoReg.FindAllStringSubmatch(body, -1)
	for _, v := range headInfo {
		student.Day = v[1]
		student.Semester = v[2]
		student.StudentName = v[4]
	}
	fmt.Println(student)
	database.G_db.Create(&student)

	tableReg := regexp.MustCompile(CLASSTABLE)
	table := tableReg.FindAllStringSubmatch(body, -1)

	classsReg := regexp.MustCompile(CLASSINFO)
	classs := classsReg.FindAllStringSubmatch(table[0][0], -1)

	for _, c := range classs {
		class.Teacher = c[1]
		class.Day = c[3]
		class.Lesson = c[4]
		class.Location = c[6]
		class.RawWeek = c[5]
	}
	fmt.Println(class)
	database.G_db.Create(&class)
	ch <- 1
}

// 获取每页的 body 信息
func GetBody(url string) string {
	userAgent := `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36`
	c := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", userAgent)
	resp, err := c.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Println("Failed to get the website information")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	return string(body)
}
