package urlshort

import "net/http"

// Returns a http.HandlerFunc that maps paths to their corresponding URL (key:val)
// the fallback http.Handler is called when path doesn't exist
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if pathsToUrls[r.URL.Path] != "" {
			http.Redirect(w, r, pathsToUrls[r.URL.Path], http.StatusFound)
		}
		fallback.ServeHTTP(w, r)
	})
}
