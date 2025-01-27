# Universal Media Downloader

A simple and user-friendly tool for downloading videos, audio, images, GIFs, and other media from various platforms. This program supports both direct file downloads and video/audio downloads using [yt-dlp](https://github.com/yt-dlp/yt-dlp). Designed for everyone to use, this downloader offers flexibility and ease of use.

---

## Features

- Supports direct downloads of images, audio, GIFs, and other files.
- Leverages `yt-dlp` to download videos and audio from supported platforms like:
  - **YouTube**
  - **Vimeo**
  - **TikTok**
  - **Facebook**
  - **Twitter (X)**
  - **Instagram**
  - **Reddit**
  - **Threads (threads.net)**
  - **Pinterest**
  - **Udemy**
  - **Coursera** and many more.
- Automatically organizes all downloads into a `downloads/` directory.

---

## Getting Started

### 1. Clone or Download the Repository
You can clone this repository or download the ZIP file from the repository page.

```bash
# Clone the repository
git clone https://github.com/NathanCordeiro/universal-media-downloader.git

# Navigate to the project directory
cd universal-media-downloader
```

Alternatively, download the ZIP file and extract it to a folder of your choice.

---

### 2. Running the Downloader

> [!NOTE]  
> If this does not work, first complete step `3` then return to step `2`.

#### Running Directly

Ensure that you have Go installed on your machine. Then, run the following command:

```bash
go run main.go
```

---

### 3. Configuration and Requirements

This program relies on `yt-dlp` and `ffmpeg` for downloading video and audio files from supported sites. Follow these steps to set them up:

#### Install yt-dlp

1. [Download yt-dlp](https://github.com/yt-dlp/yt-dlp/releases/latest) for your operating system.
2. Place the downloaded file in `C:\Users\<YourUsername>\AppData\Local\Programs\yt-dlp\` (or any preferred directory).
3. Add `yt-dlp` to your PATH (optional).
   - Open System Properties > Advanced > Environment Variables.
   - Add the directory containing `yt-dlp` (e.g., `C:\Users\<YourUsername>\AppData\Local\Programs\yt-dlp\`) to your PATH variable.

#### Install ffmpeg

1. [Download ffmpeg](https://ffmpeg.org/download.html) for your operating system.
2. Extract the files to a directory like `C:\Users\<YourUsername>\ffmpeg\bin`.
3. Add the `bin` directory containing `ffmpeg.exe` to your PATH variable using the same method as above.

---

### Customizing Supported Sites

`yt-dlp` supports thousands of sites. You can view the full list by running:

```bash
yt-dlp --list-extractors
```

To add a new site, update the section of the Go program where supported platforms are detected (lines containing `strings.Contains(urlInput, ...)`) with the new site's domain.

---

## Building an Executable

To create a standalone executable, follow these steps:

1. Ensure Go is installed on your machine.
2. In the project directory, run the following command:

```bash
go build -o UniversalDownloader.exe main.go
```

This will generate a file named `UniversalDownloader.exe` in the same directory. You can share this file with others who wish to use the downloader.

#### Using the Executable

You can now run the `.exe` directly:

1. Locate the `.exe` file.
2. Double-click it to open and start using the downloader.

---

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT). Feel free to use, modify, and distribute this software.

---

## Acknowledgments

- This downloader relies on the incredible work of the [yt-dlp](https://github.com/yt-dlp/yt-dlp) project.
- Special thanks to the developers of `ffmpeg`, a versatile and essential tool for multimedia processing.

---

## Contribution

If you'd like to contribute to this project, feel free to fork it, make your changes, and submit a pull request. Feedback, feature requests, and suggestions are always welcome!

