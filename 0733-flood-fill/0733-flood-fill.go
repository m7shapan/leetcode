func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    startedColor := image[sr][sc]
    
    if startedColor != color {
        image[sr][sc] = color
        fillIn4D(image, sr, sc, color, startedColor)
    }
 
    return image
}

func fillIn4D(image [][]int, sr int, sc int, color int, startedColor int) {
    // up    
    if sr > 0 && image[sr-1][sc] == startedColor {
        image[sr-1][sc] = color
        fillIn4D(image, sr-1, sc, color, startedColor)
    }

    // down
    if sr+1 < len(image) && image[sr+1][sc] == startedColor {
        image[sr+1][sc] = color
        fillIn4D(image, sr+1, sc, color, startedColor)
    }
    
    // left
    if sc > 0 && image[sr][sc-1] == startedColor {
        image[sr][sc-1] = color
        fillIn4D(image, sr, sc-1, color, startedColor)
    }
    
    // right
    if sc+1 < len(image[sr]) && image[sr][sc+1] == startedColor {
        image[sr][sc+1] = color
        fillIn4D(image, sr, sc+1, color, startedColor)
    }
}