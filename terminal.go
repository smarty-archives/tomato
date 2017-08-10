package tomato

import "io"

type SmartyTerminal struct {
	Reader io.Reader
	Writer io.Writer
}

func (this *SmartyTerminal) Write(p []byte) (n int, err error) {
	return this.Writer.Write(p)
}

func (this *SmartyTerminal) Read(p []byte) (n int, err error) {
	return this.Reader.Read(p)
}
