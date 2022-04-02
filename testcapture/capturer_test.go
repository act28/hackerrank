package capturer

import (
	"fmt"
	"os"
	"strings"
)

func ExampleCaptureStdout() {
	out := CaptureStdout(func() {
		fmt.Fprint(os.Stdout, "foo")
	})

	fmt.Println(strings.Join(out, "\n"))
}

func ExampleCaptureStderr() {
	out := CaptureStderr(func() {
		fmt.Fprint(os.Stderr, "bar")
	})

	fmt.Println(strings.Join(out, "\n"))
}

func ExampleCaptureOutput() {
	out := CaptureOutput(func() {
		fmt.Fprint(os.Stdout, "foo")
		fmt.Fprint(os.Stderr, "bar")
	})

	fmt.Println(strings.Join(out, "\n"))
}
