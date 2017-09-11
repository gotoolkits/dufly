package common

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gotoolkits/backend/config"
)

func CommonandLs(path string) string {

	cmd := exec.Command("/bin/ls", "-l", path)
	d, err := ExecuteCmd(cmd)

	if err != nil {
		fmt.Println("execute the 'ls' command failed!", err)
		//os.Exit(1)
	}

	return string(d)

}

func GetFilelist(path string) []string {
	var paths []string
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		paths = append(paths, path)
		return nil
	})

	if err != nil {
		fmt.Errorf("filepath returned %v\n", err)
		return nil
	}

	return paths
}

func Execute(c *exec.Cmd) ([]byte, error) {
	var err error
	var d []byte

	stdout, _ := c.StdoutPipe()
	err = c.Start()
	if err != nil {
		return nil, err
	}

	d, err = ioutil.ReadAll(stdout)

	end := len(d) - 2
	js := d[1:end]

	if err != nil {
		return nil, err
	}

	err = c.Wait()
	if err != nil {
		return nil, err
	}

	return js, nil

}

func ExecuteCmd(c *exec.Cmd) ([]byte, error) {
	var err error
	var d []byte

	stdout, err := c.StdoutPipe()
	if err != nil {
		return nil, err
	}
	defer stdout.Close()

	if err := c.Start(); err != nil {
		return nil, err
	}

	d, err = ioutil.ReadAll(stdout)

	//fmt.Println(d)
	//end := len(d) - 2
	//js := d[1:end]

	if err != nil {
		return nil, err
	}

	err = c.Wait()
	if err != nil {
		return nil, err
	}

	return d, nil

}

func GetValueFromBytes(bs []byte, jkey string) (string, error) {
	jc, err := rrconfig.LoadJsonConfigFromBytes(bs)
	if err != nil {
		return "", err
	}

	v, err := jc.GetString(jkey)

	if err != nil {
		return "", err
	}

	return v, nil
}

func GetValueFromFile(path string, jkey string) (string, error) {

	jc, err := rrconfig.LoadJsonConfigFromFile(path)

	if err != nil {
		return "", err
	}

	v, err := jc.GetString(jkey)

	if err != nil {
		return "", err
	}

	return v, nil

}

func GetSliceFromFile(path string, jkey string) ([]string, error) {

	jc, err := rrconfig.LoadJsonConfigFromFile(path)

	if err != nil {
		return nil, err
	}

	v, err := jc.GetStringSlice(jkey)

	if err != nil {
		return nil, err
	}

	return v, nil

}

func GetFileToArry(path string) ([]string, error) {

	fb, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	strs := strings.Split(string(fb), "\n")

	return strs, nil

}

func GetFileToStr(path string) (string, error) {

	fb, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(fb), nil

}
