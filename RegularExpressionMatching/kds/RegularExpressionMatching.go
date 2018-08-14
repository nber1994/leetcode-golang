package kds

/**
 https://segmentfault.com/a/1190000009583575
 */
func IsMatch(s string, p string) bool {
	pLen, sLen := len(p), len(s)
	if 0 == pLen {
		return 0 == sLen
	}
	if 1 == pLen {
		return 1 == sLen && (s[0] == p[0] || '.' == p[0])
	}

	if '*' == p[1] {
		for 0 != sLen && (s[0] == p[0] || '.' == p[0]) {
			if IsMatch(s, string([]byte(p)[2:])) {
				return true
			}
			s = string([]byte(s)[1:])
			sLen--
		}
		return IsMatch(s, string([]byte(p)[2:]))
	} else {
		if 0 == sLen || (p[0] != s[0] && '.' != p[0]) {
			return false
		} else {
			return IsMatch(string([]byte(s)[1:]), string([]byte(p)[1:]))
		}
	}
}
