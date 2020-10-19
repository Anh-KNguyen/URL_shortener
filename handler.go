package urlshort

import "net/http"

// MapHandler returns a http.HandlerFunc that will map any paths to their corresponding URL (key:val)
// When the path doesn't exist, the fallback http.Handler will be called instead
func MapHandler(pathsToUrls map[string] string, fallback http.Handler) http.HandlerFunc {
	return nil
}