package main

func main() {
	//initialize an array of arrays of ints of n size
	size := 4
	iArray := createSquareArray(size)

	//initialize variables to be used as move counters & positions
	rightMoves := size
	leftMoves := size - 1
	currentL := 0
	currentR := 0

	for i := 0; i < size; i++ {
		//Move right on the snake and pass back the counter and current positions
		currentR, currentL, rightMoves = moveRight(currentR, currentL, rightMoves, iArray)

		//Move down on the snake and pass back the counters and current positions
		currentR, currentL, leftMoves = moveDown(currentR, currentL, leftMoves, iArray)
		
		//Move left on the snake and pass back the counters and current positions
		currentR, currentL, rightMoves = moveLeft(currentR, currentL, rightMoves, iArray)
		
		//Move up on the snake and pass back the counters and current positions
		currentR, currentL, leftMoves = moveUp(currentR, currentL, leftMoves, iArray)
	}

}

func createSquareArray(size int) (iArray [][]int) {
	counter := 1
	nArray := []int{}
	for i := 1; i <= size; i++ {
		nArray = []int{}
		for i := 1; i <= size; i++ {
			nArray = append(nArray, (counter))
			counter++
		}
		iArray = append(iArray, nArray)
	}
	return iArray
}

func moveRight(currentR int, currentL int, moves int, iAr [][]int) (int, int, int) {
	for i := 0; i < moves; i++ {
		println(iAr[currentL][currentR])
		currentR++
	}
	moves--
	return currentR - 1, currentL + 1, moves
}

func moveDown(currentR int, currentL int, moves int, iAr [][]int) (int, int, int) {
	for i := 0; i < moves; i++ {
		println(iAr[currentL][currentR])
		currentL++
	}
	moves--
	return currentR - 1, currentL - 1, moves
}

func moveLeft(currentR int, currentL int, moves int, iAr [][]int) (int, int, int) {
	for i := 0; i < moves; i++ {
		println(iAr[currentL][currentR])
		currentR--
	}
	moves--
	return currentR + 1, currentL - 1, moves
}

func moveUp(currentR int, currentL int, moves int, iAr [][]int) (int, int, int) {
	for i := 0; i < moves; i++ {
		println(iAr[currentL][currentR])
		currentL--
	}
	moves--
	return currentR + 1, currentL + 1, moves
}
