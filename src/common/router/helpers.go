package router

import "strings"

func handleBackslashAtStart(url string, shouldContain bool) string {
	if strings.HasPrefix(url, "/") && shouldContain {
		return url
	}

	if !shouldContain {
		return strings.TrimPrefix(url, "/")
	}

	return "/" + url
}

func handleBackslashAtEnd(url string, shouldContain bool) string {
	if strings.HasSuffix(url, "/") && shouldContain {
		return url
	}

	if !shouldContain {
		return strings.TrimSuffix(url, "/")
	}

	return url + "/"
}

func handlePath(path string) string {
	path = handleBackslashAtStart(path, true)
	return path
}
