func maxAreaOfIsland(grid [][]int) int {
    maxArea := 0

    for i:=0;i<len(grid);i++ {
        for j:=0;j<len(grid[i]);j++{
            var area int
            if grid[i][j] == 1 {
                grid[i][j] = 2

                area = areOf4D(grid, i, j, 1)


                if area > maxArea {
                    maxArea = area
                }
            }

        }
    }

    return maxArea
}

func areOf4D(grid [][]int, x, y int, area int) int {
    // up
    if x > 0 && grid[x-1][y] == 1 {
        grid[x-1][y] = 2
        area++
        area = areOf4D(grid, x-1, y, area)
    }
    // right
    if y+1 < len(grid[x]) && grid[x][y+1] == 1 {
        grid[x][y+1] = 2
        area++
        area = areOf4D(grid, x, y+1, area)
    }
    // down
    if x+1 < len(grid) && grid[x+1][y] == 1 {
        grid[x+1][y] = 2
        area++
        area = areOf4D(grid, x+1, y, area)
    }
    // left
    if y > 0 && grid[x][y-1] == 1 {
        grid[x][y-1] = 2
        area++
        area = areOf4D(grid, x, y-1, area)
    }

    return area
}