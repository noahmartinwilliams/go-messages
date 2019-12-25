package main

func escapeString(input string) string {
	ret := ""
	for x := 0 ; x < len(input) ; x ++ {
		if input[x] == '<' {
			ret = ret + "&lt;"
		} else if input[x] == '>' {
			ret = ret + "&gt;"
		} else if input[x] == '&' {
			ret = ret + "&amp;"
		} else {
			ret = ret + string(input[x])
		}
	}
	return ret
}

func MsgEscape(inputc chan Message) chan Message {
	retc := make(chan Message, 10)
	go func() {
		defer close(retc)
		for input := range(inputc) {
			retc <- Message{Name:escapeString(input.Name), Contents:escapeString(input.Contents)}
		}
	} ()
	return retc
}

func MsgToHTML(inputc chan Message) chan string {
	retc := make(chan string, 10)
	go func() {
		defer close(retc)
		for input := range(inputc) {
			str := "<div style=\"background-color:green\"><h3>" + input.Name + "</h3><p>" + input.Contents + "</p></div>"
			retc <- str
		}
	} ()
	return retc
}
