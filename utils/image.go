package utils

import (
	"bytes"
	"io"
	"mime/multipart"

	"github.com/disintegration/imaging"
)

func ResizeAndCompress(file multipart.File) (*bytes.Buffer, error) {
	// Seek to beginning in case file was read before
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return nil, err
	}

	img, err := imaging.Decode(file)
	if err != nil {
		return nil, err
	}

	// Resize (max width 1024)
	img = imaging.Resize(img, 1024, 0, imaging.Lanczos)

	buf := new(bytes.Buffer)
	err = imaging.Encode(buf, img, imaging.JPEG, imaging.JPEGQuality(80))
	if err != nil {
		return nil, err
	}

	return buf, nil
}
