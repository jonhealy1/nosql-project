package main

type freelist struct {
	maxPage       pgnum   // Holds the maximum page allocated. maxPage*PageSize = fileSize
	releasedPages []pgnum // Pages that were previouslly allocated but are now free
}

func newFreelist() *freelist {
	return &freelist{
		maxPage:       initialPage,
		releasedPages: []pgnum{},
	}
}

func (fr *freelist) getNextPage() pgnum {
	// If possible, fetch pages first from the released pages.
	// Else, increase the maximum page
	if len(fr.releasedPages) != 0 {
		pageID := fr.releasedPages[len(fr.releasedPages)-1]
		fr.releasedPages = fr.releasedPages[:len(fr.releasedPages)-1]
		return pageID
	}
	fr.maxPage += 1
	return fr.maxPage
}

func (fr *freelist) releasePage(page pgnum) {
	fr.releasedPages = append(fr.releasedPages, page)
}
