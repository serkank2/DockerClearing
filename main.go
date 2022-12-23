package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	DeleteContainer()
	DeleteVolume()
	DeleteImage()
}

func DeleteVolume() {
	cmd := exec.Command("docker", "volume", "ls", "--format", "{{.Name}}")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	container := strings.Split(out.String(), "\n")
	fmt.Println(container)
	for _, c := range container {
		err := exec.Command("docker", "volume", "rm", "-f", c).Run()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func DeleteContainer() {
	cmd := exec.Command("docker", "container", "ls", "--format", "{{.ID}}")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	container := strings.Split(out.String(), "\n")
	fmt.Println(container)
	for _, c := range container {
		err := exec.Command("docker", "rm", "-f", "container", c).Run()
		if err != nil {
			fmt.Println(err)
		}
	}
}
func DeleteImage() {
	cmd := exec.Command("docker", "image", "ls", "--format", "{{.ID}}")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	container := strings.Split(out.String(), "\n")
	fmt.Println(container)
	for _, c := range container {
		err := exec.Command("docker", "image", "rm", "-f", c).Run()
		if err != nil {
			fmt.Println(err)
		}
	}
}
