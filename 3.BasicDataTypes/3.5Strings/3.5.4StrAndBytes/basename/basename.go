// basename removes directory components and a .suffix.
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
package basename

//The ﬁrs t versiono f basename do es al l thewor k withoutthe help of librar ies
func Basename1(str string) string {
	//Discard last '/' and everything before.
	for i := len(str); i >= 0; i-- {
		if str[i] == '/' {
			//this make substring object which point to the underlaying
			//memory of the original string. No new memory is allocated
			str = str[i+1:]
			break
		}
	}
	// Preserve everything befor last '.'.
	for i := len(str); i >= 0; i-- {
		if str[i] == '.' {
			str = str[:i]
			break
		}
	}
	return str
}
