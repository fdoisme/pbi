package utils

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
)

func UploadFiles(ctx *fiber.Ctx, files []*multipart.FileHeader) ([]string, error) {

	scheme := ctx.Get("X-Forwarded-Proto")
	if scheme == "" {
		if ctx.Context().IsTLS() {
			scheme = "https"
		} else {
			scheme = "http"
		}
	}
	host := ctx.Hostname()

	uploadDir := "uploads"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return []string{}, err
	}
	var result []string
	for i, file := range files {
		file.Filename = fmt.Sprintf("%d%s", time.Now().Unix(), filepath.Ext(file.Filename))
		filePath := filepath.Join(uploadDir, fmt.Sprintf("%d%s", i, file.Filename))
		if err := ctx.SaveFile(file, filePath); err != nil {
			return []string{}, err
		}
		// fmt.Printf("File %s saved to %s\n", file.Filename, filePath)
		// fmt.Printf("[%s://%s/%s/%s]\n", scheme, host, uploadDir, file.Filename)
		result = append(result, fmt.Sprintf("%s://%s/%s/%d%s", scheme, host, uploadDir, i, file.Filename))
	}
	return result, nil
}
