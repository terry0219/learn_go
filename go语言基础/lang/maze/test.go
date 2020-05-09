package main

import (
	"fmt"
	"os"
)

/*
6代表row, 5代表col
首先0 1 0 0 0应该是一个slice，[0 1 0 0 0]; 下面几行也是
[0 1 0 0 0],[0 0 0 1 0],[0 1 0 1 0]... 
最后在把这些slice组成到一个slice里
[[0 1 0 0 0],[0 0 0 1 0],[0 1 0 1 0]...]


6 5
0 1 0 0 0
0 0 0 1 0
0 1 0 1 0
1 1 1 0 0
0 1 0 0 1
0 1 0 0 0
 
   -1  0 1 2 3 4
 -1   |- - - - - -        
  0   |0 1 0 0 0
  1   |0 0 0 1 0
  2   |0 1 0 1 0
  3   |1 1 1 0 0
  4   |0 1 0 0 1
  5   |0 1 0 0 0

以左上角的0开始走，按照上左下右的顺序;
1. 0的上面的坐标{-1 0}
2. 0的左边的坐标{0 -1}
3. 0的下面的坐标{1, 0}
4. 0的右边的坐标{0, 1}
*/

//生成二维数组(迷宫)
func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil{
		panic(err)
	}

	var row,col int
	fmt.Fscanf(file,"%d %d",&row,&col) //注意要用&row取地址，因为Fscanf函数会改变row col值
	//fmt.Println(row,col) // 6 5
	maze := make([][]int, row)  //二维数组, 一步步分析; 先是一个slice[],这个slice里面是[]int类型，长度是row
	//fmt.Println(maze) //[[] [] [] [] [] []]
	for i:= range maze{
		maze[i] = make([]int,col) //初始化里面的slice[[0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0]]
		for j:= range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	//fmt.Println(maze) //[[0 1 0 0 0] [0 0 0 1 0] [0 1 0 1 0] [1 1 1 0 0] [0 1 0 0 1] [0 1 0 0 0]]
	return maze 
}

//定义坐标 i代表row  j代表col, 从0开始算起; 初始坐标就是point{0,0}
type point struct {
	i,j int
}

//按照上左下右的顺序 坐标值
var dirs = [4]point{
	{-1,0},{0,-1},{1,0},{0,1}
}

func walk(maze [][]int, start, end point) {

}


/*
写代码技巧
1. 先把walk函数的参数定义好
	func walk([][]int, start,end point) {

	}
2. 在main函数中调用walk
	walk(maze, point{0,0}, point{6,5})
3. 最后在具体是先walk里面得逻辑
*/

//走迷宫
func walk([][]int, start, end point){
	//还需要维护一个二维数组，记录从start到end，一共走了多少步;最后的路径就是由这个建出来的
	//先初始化二维数组steps
	steps := make([][]int, len(maze))
	for i := range steps{
		steps[i] = make([]int, len(maze[i]))
	}

	//把起点加进队列里面
	Q := []point{start}
	
}


func main() {
	maze := readMaze("./maze.in")
	for _,row := range maze{
		for _,val := range row{
			fmt.Printf("%d ",val)
			/*
			生成迷宫
			0 1 0 0 0
			0 0 0 1 0
			0 1 0 1 0
			1 1 1 0 0
			0 1 0 0 1
			0 1 0 0 0
			*/
		}
		fmt.Println() //换行
	}

	//生成迷宫后，就要尝试走迷宫了
	//len(maze) - 1 == 5  and  len(maze[0] - 1) == 4
	walk(maze, point{0,0}, point{len(maze)-1, len(maze[0]-1)})

}
