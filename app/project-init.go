package app

import (
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/go-getter"
)

const templatesSource = "github.com/dansc11/sls-tf//templates"

func InitProject(path string) error {
	fmt.Println("Initialising sls-tf project...")

	fmt.Println("Copying template files...")
	if err := downloadTemplateFiles(path); err != nil {
		log.Fatal("DOWNLOAD ERROR: " + err.Error())
		return err
	}

	return nil
}

func downloadTemplateFiles(destPath string) error {
	if destPath == "" {
		pwd, err := os.Getwd()

		if err != nil {
			return err
		}

		destPath = pwd
	}

	return getter.GetAny(destPath, templatesSource)
}
