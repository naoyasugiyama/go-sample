package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// TODO: Add a Read([]byte) (int, error) method to MyReader.
type MyReader struct{}

/*
ASCII文字 'A' の無限ストリームを出力する Reader 型を実装してください。

func (e *MyReader) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}*/

func (a MyReader) Read(rb []byte) (n int, err error) {
	for n, err = 0, nil; n < len(rb); n++ {
		rb[n] = 'A'
	}
	return
}

// rot 13 Reader
type rot13Reader struct {
	r io.Reader
}

func rot13(c byte) byte {
	switch {
	case 'A' <= c && c <= 'Z':
		return (c-'A'+13)%26 + 'A'
	case 'a' <= c && c <= 'z':
		return (c-'a'+13)%26 + 'a'
	default:
		return c
	}
}

func (r *rot13Reader) Read(rb []byte) (n int, err error) {
	n, err = r.r.Read(rb)

	if err != nil {
		return n, err
	}

	for i := range rb {
		rb[i] = rot13(rb[i])
	}

	return n, nil
}

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)

	for {
		n, err := r.Read(b)
		fmt.Printf("n=%v err=%v b=%v\n", n, err, b)
		fmt.Printf("b[:n]=%q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	fmt.Println("-----------------------------------")

	// reader
	//	reader.Validate(MyReader{})
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r0 := rot13Reader{s}
	io.Copy(os.Stdout, &r0)
	/*
		r1 := rot13Reader{r0}
		io.Copy(os.Stdout, &r1)
	*/
}
