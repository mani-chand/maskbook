// --- HELPER FUNCTIONS ---
package utils

import (
	"encoding/base64"
	"errors"
	"strings"
)

// getMimeType parses the data URL (e.g., "data:image/png;base64,...")
// and returns the MIME type (e.g., "image/png")
func GetMimeType(dataURL string) (string, error) {
	// Find "data:"
	startIndex := strings.Index(dataURL, "data:")
	if startIndex == -1 {
		return "", errors.New("invalid data URL: missing 'data:' prefix")
	}
	// Find ";base64,"
	endIndex := strings.Index(dataURL, ";base64,")
	if endIndex == -1 {
		return "", errors.New("invalid data URL: missing ';base64,' separator")
	}

	// The MIME type is between "data:" and ";base64,"
	mimeType := dataURL[startIndex+5 : endIndex]
	return mimeType, nil
}

// mimeTypeToExtension maps a MIME type to a file extension.
// Add any other file types you want to support here.
func MimeTypeToExtension(mimeType string) (string, error) {
	switch mimeType {
	case "image/jpeg":
		return ".jpg", nil
	case "image/png":
		return ".png", nil
	case "image/gif":
		return ".gif", nil
	case "video/mp4":
		return ".mp4", nil
	case "video/quicktime":
		return ".mov", nil
	default:
		return "", errors.New("unsupported file type: " + mimeType)
	}
}

// decodeBase64File is a helper function to strip the prefix
// (e.g., "data:image/png;base64,") and decode the data.
func DecodeBase64File(dataURL string) ([]byte, error) {
	// Find the comma
	commaIndex := strings.Index(dataURL, ",")
	if commaIndex == -1 {
		return nil, errors.New("invalid base64 data URL: missing comma")
	}

	// Get the part after the comma
	base64Data := dataURL[commaIndex+1:]

	// Decode the string
	decoded, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return nil, errors.New("failed to decode base64 data")
	}
	return decoded, nil
}
