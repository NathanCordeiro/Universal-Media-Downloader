package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

// Sanitize filenames to remove invalid characters.
func sanitizeFileName(fileName string) string {
	reg := regexp.MustCompile(`[<>:"/\\|?*]+`)
	return reg.ReplaceAllString(fileName, "_")
}

// Get the filename from URL if Content-Disposition is absent.
func getFileNameFromURL(fileURL string) string {
	segments := strings.Split(fileURL, "/")
	fileName := sanitizeFileName(segments[len(segments)-1])
	if fileName == "" {
		fileName = "downloaded_file"
	}
	return fileName
}

// Get filename from Content-Disposition header.
func getFileNameFromHeader(header http.Header, defaultName string) string {
	contentDisposition := header.Get("Content-Disposition")
	if strings.Contains(contentDisposition, "filename=") {
		parts := strings.Split(contentDisposition, "filename=")
		fileName := strings.Trim(parts[len(parts)-1], "\"")
		return sanitizeFileName(fileName)
	}
	return defaultName
}

// Download files like images, GIFs, and audio directly.
func downloadDirect(url, filePath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Create the file.
	out, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer out.Close()

	// Copy data to the file.
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("error saving file: %v", err)
	}

	return nil
}

// Download videos or audio using yt-dlp.
func downloadWithYTDLP(url string) error {
	cmd := exec.Command("C:\\Users\\Nathan\\AppData\\Local\\Programs\\yt-dlp\\yt-dlp.exe", url, "-o", "downloads/%(title)s.%(ext)s", "--ffmpeg-location", "C:\\Users\\Nathan\\ffmpeg\\bin")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	fmt.Println("=== Universal Media Downloader ===")
	fmt.Println("Type 'exit' to quit")

	// Create downloads directory.
	if err := os.MkdirAll("downloads", 0755); err != nil {
		fmt.Printf("Error creating downloads directory: %v\n", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nEnter URL to download: ")
		urlInput, _ := reader.ReadString('\n')
		urlInput = strings.TrimSpace(urlInput)

		if urlInput == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		if urlInput == "" {
			continue
		}

		// Detect YouTube or other video sites and delegate to yt-dlp.
		if strings.Contains(urlInput, "youtube.com") ||
			strings.Contains(urlInput, "youtu.be") ||
			strings.Contains(urlInput, "vimeo.com") ||
			strings.Contains(urlInput, "dailymotion.com") ||
			strings.Contains(urlInput, "tiktok.com") ||
			strings.Contains(urlInput, "facebook.com") ||
			strings.Contains(urlInput, "twitter.com") ||
			strings.Contains(urlInput, "instagram.com") ||
			strings.Contains(urlInput, "threads.net") || // Threads uses threads.net
			strings.Contains(urlInput, "reddit.com") ||
			strings.Contains(urlInput, "linkedin.com") ||
			strings.Contains(urlInput, "pinterest.com") ||
			strings.Contains(urlInput, "udemy.com") ||
			strings.Contains(urlInput, "coursera.org") {
			fmt.Println("Downloading video via yt-dlp...")
			if err := downloadWithYTDLP(urlInput); err != nil {
				fmt.Printf("Error downloading with yt-dlp: %v\n", err)
			} else {
				fmt.Println("Video downloaded successfully.")
			}
			continue
		}

		// Try direct download for images, audio, GIFs, etc.
		resp, err := http.Head(urlInput) // Use HEAD to check the type.
		if err != nil {
			fmt.Printf("Error checking URL: %v\n", err)
			continue
		}

		contentType := resp.Header.Get("Content-Type")
		fileName := getFileNameFromHeader(resp.Header, getFileNameFromURL(urlInput))
		filePath := filepath.Join("downloads", fileName)

		// Check supported content types.
		switch {
		case strings.HasPrefix(contentType, "image/"), strings.HasPrefix(contentType, "audio/"), strings.Contains(contentType, "gif"):
			fmt.Printf("Downloading file: %s\n", fileName)
			if err := downloadDirect(urlInput, filePath); err != nil {
				fmt.Printf("Error downloading file: %v\n", err)
			} else {
				fmt.Println("File downloaded successfully.")
			}

		default:
			fmt.Printf("Unsupported content type: %s\n", contentType)
		}
	}
}
