package FreiOrderBot

import "strconv"

func TreeToPlain(tree OrderTree) *map[int]Goods {
	var tempMap = make(map[int]Goods, 100)

	for _, i := range tree.Cons.Bar.Cups {
		tempMap[i.Id] = i
	}

	for _, i := range tree.Cons.Bar.Others {
		tempMap[i.Id] = i
	}

	for _, i := range tree.Cons.Chemical {
		tempMap[i.Id] = i
	}

	for _, i := range tree.Packages {
		tempMap[i.Id] = i
	}

	for _, i := range tree.AlContainer {
		tempMap[i.Id] = i
	}

	for _, i := range tree.PContainer {
		tempMap[i.Id] = i
	}

	for _, i := range tree.Factory {
		tempMap[i.Id] = i
	}

	return &tempMap
}

func InsertCounter(text string, count int) string {
	res := ""

	for i := len(text) - 1; i >= 0; i-- {
		if text[i] == '-' {
			res = text[:i-1] + " - " + strconv.Itoa(count) + " " + text[i+1:] + "\n"
			i = 1
		}
	}

	return res
}
