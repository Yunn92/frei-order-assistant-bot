package FreiOrderBot

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
