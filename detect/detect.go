package detect

import "github.com/mouselog/mouselog/trace"

func checkBot(t *trace.Trace) (int, string, int, int, int) {
	isBot, reason, rule, start, end := checkStraightLine(t)
	if isBot != 0 {
		return isBot, reason, rule, start, end
	}

	return 0, "", RuleNone, -1, -1
}

func CheckBot(t *trace.Trace) (int, string, int, int, int) {
	isBot, reason, rule, start, end := checkBot(t)
	t.IsBot = isBot
	t.Reason = reason
	t.RuleId = rule
	t.RuleStart = start
	t.RuleEnd = end

	return isBot, reason, rule, start, end
}

func GetDetectResult(ss *trace.Session, url string) *trace.Session {
	t, ok := ss.UrlMap[url]
	if url != "" || ok {
		CheckBot(t)
	}

	return ss
}