package block

func CalculateNumBlocks(scrWidth int, scrHeight int, blkSizeX int, blkSizeY int) (numBlkX int, numBlkY int) {
	numBlkX = scrWidth / blkSizeX
	numBlkY = scrHeight / blkSizeY
	return
}