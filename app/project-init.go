package app

import (
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/go-getter"
)

const templatesSource = "github.com/dansc11/sls-tf//templates"

func InitProject() error {
	fmt.Println("Initialising sls-tf project...")

	fmt.Println("Copying template files...")
	if err := downloadTemplateFiles(); err != nil {
		log.Fatal("DOWNLOAD ERROR: " + err.Error())
		return err
	}

	return nil
}

func downloadTemplateFiles() error {
	pwd, err := os.Getwd()

	if err != nil {
		return err
	}

	return getter.GetAny(pwd, templatesSource)
}
