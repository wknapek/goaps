package solar

import (
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func downloadContent(url string, file string) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("download error from URL: %s", url)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading content %s: %v\n", url, err)
	}
	resp.Body.Close()
	content := string(b[:])
	removeEmptyLines(&content)
	fmt.Println(content + "\n")
	fmt.Println("=========================================================")
	count := 0
	WordCount(content, &count)
	maxSentenceWords(content)
	fmt.Println(count)
	_ = file
}

func maxSentenceWords(s string) {
	dots := make([]string, 0)
	ques := make([]string, 0)
	exs := make([]string, 0)
	news := make([]string, 0)
	idx := findFirstseparator(s)
	var numberAllSentences int
	var totalwordslenth int
	var totalwordsnumber int
	//fmt.Printf("index======>%d\n", idx)
	for len(s) != 0 && idx > 0 && idx < len(s) {
		tmprntens := s[:idx]
		s = s[idx+1:]
		if string(tmprntens[len(tmprntens)-1]) == "." {
			dots = append(dots, tmprntens)
			//fmt.Println(tmprntens)
			numberAllSentences++
		}
		if string(tmprntens[len(tmprntens)-1]) == "?" {
			ques = append(ques, tmprntens)
			//fmt.Println(tmprntens)
			numberAllSentences++
		}
		if string(tmprntens[len(tmprntens)-1]) == "!" {
			exs = append(exs, tmprntens)
			//fmt.Println(tmprntens)
			numberAllSentences++
		}
		if string(tmprntens[:len(tmprntens)-1]) == "\\n" {
			news = append(news, tmprntens)
			//fmt.Println(tmprntens)
			numberAllSentences++
		}
		idx = findFirstseparator(s)
	}
	for _, sen := range dots {
		totalwordsnumber += WordCountPerSentence(sen, &totalwordslenth)
	}

	for _, sen := range ques {
		totalwordsnumber += WordCountPerSentence(sen, &totalwordslenth)
	}

	for _, sen := range exs {
		totalwordsnumber += WordCountPerSentence(sen, &totalwordslenth)
	}

	for _, sen := range news {
		totalwordsnumber += WordCountPerSentence(sen, &totalwordslenth)
	}
	fmt.Printf("number of all sentenses are: %d\n", numberAllSentences)
	fmt.Printf("total words number: %d\n", totalwordsnumber)
	fmt.Printf("total words lenth: %d\n", totalwordslenth)
	fmt.Printf("avarage world lenhth is: %.2f", float32(totalwordslenth/totalwordsnumber))
}

func findFirstseparator(s string) int {
	ret := -1
	ret = 0
	dotIdx := strings.Index(s, ".")
	quesIdx := strings.Index(s, "?")
	exIdx := strings.Index(s, "!")
	newIdx := strings.Index(s, "\\n")
	if dotIdx < 0 {
		dotIdx = math.MaxInt32
	}
	if quesIdx < 0 {
		quesIdx = math.MaxInt32
	}
	if exIdx < 0 {
		exIdx = math.MaxInt32
	}
	if newIdx < 0 {
		newIdx = math.MaxInt32
	}
	/*fmt.Printf("dotIdx %d\n", dotIdx)
	fmt.Printf("quesIdx %d\n", quesIdx)
	fmt.Printf("exIdx %d\n", exIdx)
	fmt.Printf("newIdx %d\n", newIdx)*/
/*if dotIdx < quesIdx && dotIdx < exIdx && dotIdx < newIdx && dotIdx > 0 {
		ret = dotIdx + 1
	}
	if quesIdx < dotIdx && quesIdx < exIdx && quesIdx < newIdx && quesIdx > 0 {
		ret = quesIdx + 1
	}
	if exIdx < dotIdx && exIdx < quesIdx && exIdx < newIdx && exIdx > 0 {
		ret = exIdx + 1
	}
	if newIdx < dotIdx && newIdx < quesIdx && newIdx < exIdx && newIdx > 0 {
		ret = newIdx + 1
	}
	if newIdx == math.MaxInt32 && exIdx == math.MaxInt32 && quesIdx == math.MaxInt32 && dotIdx == math.MaxInt32 {
		fmt.Println("no more sentenses")
		return -1
	}
	//fmt.Printf("findFirstseparator ret %d\n", ret)
	return ret
}

func WordCountPerSentence(s string, lenofwords *int) int {
	words := strings.Fields(s)
	counter := 0
	for iter := 0; iter < len(words); iter++ {
		counter++
		*lenofwords += len(words[iter])
		//fmt.Println("word " + words[iter])
		//fmt.Printf("len of word: %d\n", len(words[iter]))
	}
	//fmt.Printf("%s ** nb. of words %d\n", s, counter)
	return counter
}

func removeEmptyLines(s *string) {
	regex, err := regexp.Compile("\n\n")
	if err != nil {
		return
	}
	*s = regex.ReplaceAllString(*s, "\n")
}

func WordCount(s string, count *int) {
	words := strings.Fields(s)
	*count = len(words)

}
