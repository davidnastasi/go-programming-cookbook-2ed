package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/davidnastasi/go-programming-cookbook-2ed/chapter2/envvar"
)

// Config ...
type Config struct {
	Version string `json:"version" required:"true"`
	IsSafe  bool   `json:"is_safe" default:"true"`
	Secret  string `json:"secret"`
}

func main() {
	var err error

	tf, err := ioutil.TempFile("", "tmp")
	if err != nil {
		panic(err)
	}
	defer tf.Close()
	defer os.Remove(tf.Name())

	secrets := `{"secret": "so so secret"}`

	_, err = tf.Write(bytes.NewBufferString(secrets).Bytes())
	//_, err = tf.Write([]byte(secrets))
	if err != nil {
		panic(err)
	}

	fmt.Println("secrets file contains =", secrets)

	if err = os.Setenv("EXAMPLE_VERSION", "1.0.0"); err != nil {
		panic(err)
	}

	if err = os.Setenv("EXAMPLE_ISSAFE", "false"); err != nil {
		panic(err)
	}

	c := Config{}
	if err = envvar.LoadConfig(tf.Name(), "EXAMPLE", &c); err != nil {
		panic(err)
	}

	fmt.Println("EXAMPLE_VERSION = ", os.Getenv("EXAMPLE_VERSION"))
	fmt.Println("EXAMPLE_ISSAFE = ", os.Getenv("EXAMPLE_ISSAFE"))

	fmt.Printf("Final Config: %#v\n", c)

}
