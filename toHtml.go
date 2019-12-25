package main

func escapeString(input string) string {
	ret := ""
	for x := 0 ; x < len(input) ; x ++ {
		if input[x] == '<' {
			ret = ret + "&lt;"
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
