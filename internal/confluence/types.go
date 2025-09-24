package confluence

type CreateConfluencePage struct {
	Title    string
	SpaceKey string
	Body     string
}

type UpdateConfluencePage struct {
	PageID  string
	Title   string
	Body    string
	Version int
}

type GetSpacePages struct {
	SpaceKey string
	Limit    int
	Start    int
}
