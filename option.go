package bmsdb

type QueryContext struct {
	path  string
	after *int64
}

func NewQueryContext(path string) *QueryContext {
	return &QueryContext{
		path: path,
	}
}

func (q *QueryContext) After(timestamp int64) *QueryContext {
	q.after = &timestamp
	return q
}

type ScanContext struct {
	path string
}

func NewScanContext(path string) *ScanContext {
	return &ScanContext{
		path: path,
	}
}
