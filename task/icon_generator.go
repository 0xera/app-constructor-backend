package task

import (
	"app-constructor-backend/task/pb"
	"errors"
	"github.com/nfnt/resize"
	"image"
	"image/png"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

type IconGenerator struct {
}

type androidIcon struct {
	directoryName string
	size          uint
}

var (
	icons = []androidIcon{
		{
			directoryName: "mipmap-mdpi",
			size:          48,
		}, {
			directoryName: "mipmap-hdpi",
			size:          72,
		}, {
			directoryName: "mipmap-xhdpi",
			size:          96,
		}, {
			directoryName: "mipmap-xxhdpi",
			size:          144,
		}, {
			directoryName: "mipmap-xxxhdpi",
			size:          192,
		},
	}
)

func (g IconGenerator) generateIcon(a *pb.App, templatesDir string) error {

	var iconProp *pb.App_Props

	for _, prop := range a.Props {
		if prop.Name == "Icon png link" {
			iconProp = prop
			break
		}
	}

	if iconProp == nil {
		return errors.New("no icon")
	}

	iconPath := filepath.FromSlash(templatesDir + "/icon.png")

	err := downloadFile(iconProp.Value, iconPath)
	if err != nil {
		return err
	}

	iconFile, err := os.Open(iconPath)
	if err != nil {
		return err
	}

	iconPng, err := png.Decode(iconFile)
	if err != nil {
		return err
	}
	err = iconFile.Close()
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	wg.Add(len(icons))

	caughtErrors := make(chan error)
	wgDone := make(chan bool)

	for _, icon := range icons {
		go rewriteIconFiles(templatesDir, iconPng, icon, &wg, caughtErrors)
	}
	go func() {
		wg.Wait()
		close(wgDone)
	}()

	select {
	case <-wgDone:
		break
	case err := <-caughtErrors:
		close(caughtErrors)
		return err
	}

	return nil
}

func rewriteIconFiles(dir string, iconPng image.Image, icon androidIcon, wg *sync.WaitGroup, caughtErrors chan error) {
	defer wg.Done()
	m := resize.Resize(icon.size, icon.size, iconPng, resize.Lanczos3)
	destFile := filepath.FromSlash(dir + "/AppConstructor/app/src/main/res/" + icon.directoryName + "/ic_launcher.png")
	out, err := os.Create(destFile)
	if err != nil {
		caughtErrors <- err
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			caughtErrors <- err
		}
	}(out)
	err = png.Encode(out, m)
	if err != nil {
		caughtErrors <- err
	}
}

func downloadFile(URL, fileName string) error {
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("image not found")
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}
