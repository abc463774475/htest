package pipe

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/abc463774475/bbtool/n_log"
	"os/exec"
	"testing"
)

func TestServer(t *testing.T) {
	cmd_go_env := exec.Command("go", "env")
	n_log.Info("info\n%v",cmd_go_env)

	stdout_env, err := cmd_go_env.StdoutPipe()
	if err != nil {
		n_log.Panic("err  %v", err)
	}

	if err = cmd_go_env.Start(); err != nil {
		n_log.Panic("err  %v", err)
	}

	data := make([]byte,1024)
	n,err := stdout_env.Read(data)
	if err != nil {
		n_log.Panic("err  %v", err)
	}

	n_log.Info("data  %v",string(data[:n]))
}

func TestClient(t *testing.T) {

	cmd_go_env := exec.Command("go", "env")
	cmd_grep := exec.Command("grep", "GOROOT")

	stdout_env, env_error := cmd_go_env.StdoutPipe()
	if env_error != nil {
		n_log.Info("Error happened about standard output pipe %v", env_error)
		return
	}

	//env_error := cmd_go_env.Start()
	if env_error := cmd_go_env.Start(); env_error != nil {
		n_log.Info("Error happened in execution %v", env_error)
		return
	}
	stdout_buf_grep := bufio.NewReader(stdout_env)

	//create input pipe for grep command
	stdin_grep, grep_error := cmd_grep.StdinPipe()
	if grep_error != nil {
		fmt.Println("Error happened about standard input pipe ", grep_error)
		return
	}

	//connect the two pipes together
	stdout_buf_grep.WriteTo(stdin_grep)

	var buf_result bytes.Buffer
	cmd_grep.Stdout = &buf_result

	//grep_error := cmd_grep.Start()
	if grep_error := cmd_grep.Start(); grep_error != nil {
		fmt.Println("Error happened in execution ", grep_error)
		return
	}

	err := stdin_grep.Close()
	if err != nil {
		fmt.Println("Error happened in closing pipe", err)
		return
	}

	//make sure all the infor in the buffer could be read
	if err := cmd_grep.Wait(); err != nil {
		fmt.Println("Error happened in Wait process")
		return
	}
	fmt.Println(buf_result.String())
}

