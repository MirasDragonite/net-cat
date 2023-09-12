package internal

import "strings"

func isValidUserName(c *Server, s string) bool {
	for _, ch := range s {
		if ch < 32 || ch > 126 {
			return false
		}
	}
	if len(s) < 3 || len(s) > 20 {
		return false
	}
	if strings.TrimSpace(s) == "" {
		return false
	}
	for _, username := range c.users {
		if s == username {
			return false
		}
	}

	return true
}

func isValidText(s string) bool {
	for _, ch := range s {
		if ch < 32 || ch > 126 {
			return false
		}
	}

	if strings.TrimSpace(s) == "" {
		return false
	}
	return true
}

func IsValidPort(port string) bool {
	for _, ch := range port {
		if ch < 48 || ch > 57 {
			return false
		}
	}
	return true
}
