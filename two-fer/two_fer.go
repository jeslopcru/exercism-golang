package twofer

import "fmt"

func ShareWith(aName string) string {

	const TWOFER = "One for %s, one for me."
	twoFerYou := "you"
	if aName != "" {
		twoFerYou = aName
	}
	return fmt.Sprintf(TWOFER, twoFerYou)
}
