package block

func CalculateBlockSize(scrWidth int, scrHeight int, nlX int, nlY int) (blkSizeX int, blkSizeY int) {
	blkSizeX = scrWidth / nlX
	blkSizeY = scrHeight / nlY
	return
}