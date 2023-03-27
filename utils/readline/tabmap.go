package readline

import (
	runewidth "github.com/mattn/go-runewidth"
)

func (rl *Instance) initTabMap() {
	rl.tabMutex.Lock()
	defer rl.tabMutex.Unlock()

	var suggestions []string
	if rl.modeTabFind {
		suggestions = rl.tfSuggestions
	} else {
		suggestions = rl.tcSuggestions
	}

	rl.tcMaxLength = 1
	for i := range suggestions {
		if rl.tcDisplayType == TabDisplayList {
			if len(rl.tcPrefix+suggestions[i]) > rl.tcMaxLength {
				rl.tcMaxLength = len([]rune(rl.tcPrefix + suggestions[i]))
			}

		} else {
			if len(rl.tcDescriptions[suggestions[i]]) > rl.tcMaxLength {
				rl.tcMaxLength = len(rl.tcDescriptions[suggestions[i]])
			}
		}
	}

	rl.tcPosX = 1
	rl.tcPosY = 1
	rl.tcOffset = 0
	rl.tcMaxX = 1

	if len(suggestions) > rl.MaxTabCompleterRows {
		rl.tcMaxY = rl.MaxTabCompleterRows
	} else {
		rl.tcMaxY = len(suggestions)
	}
}

func (rl *Instance) moveTabMapHighlight(x, y int) {
	rl.tabMutex.Lock()
	defer rl.tabMutex.Unlock()

	var suggestions []string
	if rl.modeTabFind {
		suggestions = rl.tfSuggestions
	} else {
		suggestions = rl.tcSuggestions
	}

	rl.tcPosY += x
	rl.tcPosY += y

	if rl.tcPosY < 1 {
		rl.tcPosY = 1
		rl.tcOffset--
	}

	if rl.tcPosY > rl.tcMaxY {
		rl.tcPosY--
		rl.tcOffset++
	}

	if rl.tcOffset+rl.tcPosY < 1 && len(suggestions) > 0 {
		rl.tcPosY = rl.tcMaxY
		rl.tcOffset = len(suggestions) - rl.tcMaxY
	}

	if rl.tcOffset < 0 {
		rl.tcOffset = 0
	}

	if rl.tcOffset+rl.tcPosY > len(suggestions) {
		rl.tcPosY = 1
		rl.tcOffset = 0
	}
}

func (rl *Instance) writeTabMap() {
	rl.tabMutex.Lock()
	defer rl.tabMutex.Unlock()

	var suggestions []string
	if rl.modeTabFind {
		suggestions = rl.tfSuggestions
	} else {
		suggestions = rl.tcSuggestions
	}

	if rl.termWidth < 10 {
		// terminal too small. Probably better we do nothing instead of crash
		return
	}

	maxLength := rl.tcMaxLength
	if maxLength > rl.termWidth-9 {
		maxLength = rl.termWidth - 9
	}
	maxDescWidth := rl.termWidth - maxLength - 4

	//cellWidth := strconv.Itoa(maxLength)
	//itemWidth := strconv.Itoa(maxDescWidth)

	y := 0

	moveCursorUp(1) // bit of a kludge. Really should find where the code is "\n"ing

	isTabDisplayList := rl.tcDisplayType == TabDisplayList

	var item, description string
	for i := rl.tcOffset; i < len(suggestions); i++ {
		y++
		if y > rl.tcMaxY {
			break
		}

		if isTabDisplayList {
			/*item = rl.tcPrefix + suggestions[i]
			if len(item) > maxLength {
				item = item[:maxLength-3] + "..."
			}*/
			item = runewidth.Truncate(rl.tcPrefix+suggestions[i], maxLength, "…")

			/*description = rl.tcDescriptions[suggestions[i]]
			if len(description) > maxDescWidth {
				description = description[:maxDescWidth-3] + "..."
			}*/
			description = runewidth.Truncate(rl.tcDescriptions[suggestions[i]], maxDescWidth, "…")

		} else {
			/*item = rl.tcPrefix + suggestions[i]
			if len(item) > maxDescWidth {
				item = item[:maxDescWidth-3] + "..."
			}*/
			item = runewidth.Truncate(rl.tcPrefix+suggestions[i], maxDescWidth, "…")

			/*description = rl.tcDescriptions[suggestions[i]]
			if len(description) > maxLength {
				description = description[:maxLength-3] + "..."
			}*/
			description = runewidth.Truncate(rl.tcDescriptions[suggestions[i]], maxLength, "…")
		}

		if isTabDisplayList {
			//printf("\r\n%s %-"+cellWidth+"s %s %s",
			//	highlight(rl, y), item, seqReset, description)
			printf("\r\n%s %s %s %s",
				highlight(rl, y), runewidth.FillRight(item, maxLength),
				seqReset, description)

		} else {
			//printf("\r\n %-"+cellWidth+"s %s %-"+itemWidth+"s %s",
			//	description, highlight(rl, y), item, seqReset)
			printf("\r\n %s %s %s %s",
				runewidth.FillRight(description, maxLength), highlight(rl, y),
				runewidth.FillRight(item, maxDescWidth), seqReset)
		}
	}

	if len(suggestions) < rl.tcMaxX {
		rl.tcUsedY = len(suggestions)
	} else {
		rl.tcUsedY = rl.tcMaxY
	}
}

func highlight(rl *Instance, y int) string {
	if y == rl.tcPosY {
		return seqBgWhite + seqFgBlack
	}
	return ""
}
