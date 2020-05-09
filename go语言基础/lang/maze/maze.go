/*
广度优先搜索走迷宫算法

初始迷宫，1代表墙壁
6 5
0 1 0 0 0
0 0 0 1 0
0 1 0 1 0
1 1 1 0 0
0 1 0 0 1
0 1 0 0 0

从左上角开始，走到右下角，计算一共走了多少步以及它的线路(也是一个二维数组)

1. 首先初始化二维数组[][]int，在通过读文件的方式，将数据填充到数组中，生成迷宫
2. 定义一个point结构体用于记录坐标值，定义walk(maze [][]int, start, end point)函数
3. walk函数，通过从start开始，比如是{0 0}位置开始，然后分别遍历当前坐标的上、左、下、右的位置是否满足条件
4. next坐标需要的条件: (1) 因为定义好了row col，那么next坐标就不能越界
					  (2) next坐标指向的值不能为1，因为1代表墙壁
					  (3) 
4. 如果next坐标满足条件，就把next坐标写到Q队列里，并将next坐标值加1写到steps二维数组里





1. 用循环创建二维slice
2. 使用slice来实现队列
3. 用Fscanf读取文件
4. 对Point的抽象



*/
package main

import (
	"fmt"
	"os"
)

/*
6 5
0 1 0 0 0
0 0 0 1 0
0 1 0 1 0
1 1 1 0 0
0 1 0 0 1
0 1 0 0 0
*/
//从文件中把迷宫读进来，生成一个二维数组
func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row) //先看第一个[]，是一个slice类型，它里面的元素是[]int类型; row就是slice长度，和第二个[]int没关系
	for i := range maze {
		maze[i] = make([]int, col) //先初始化好二维数组，才可以用
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j]) // 从文件中读取后，在写到二维数组里
		}
	}

	return maze
}

//点的结构, 坐标
type point struct {
	i, j int
}

//定义上左下右 坐标的位置; 后面会根据当前坐标位置，去遍历它的上左下右位置
var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

//值传递 获取当前point坐标的 下一步位置
func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

//根据当前传入的point坐标， 去grid二维数组就查询对应的数值
func (p point) at(grid [][]int) (int, bool) {
	//判断row数不能越界，不能小于0或者大于len(grid)
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}

	//判断col不能越界, 不能小于0或者大于len(grid[p.i])
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

//具体走迷宫的函数, maze是迷宫图，start end是起始结束坐标，返回[][]int(具体走的路径)
func walk(maze [][]int,
	start, end point) [][]int {
	//steps也是一个二维数组，记录从开始点到结束点走的路径; 先初始化二维数组
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	//定义slice，从起始位置开始,用于队列;将下一步要遍历的point坐标压入队列
	Q := []point{start}

	//下面就开始探索了 len(Q) > 0 只要队列里有数据就遍历
	for len(Q) > 0 {
		cur := Q[0] //取出队列中第一个元素，遍历
		Q = Q[1:]   //然后把队列中的第一个元素拿掉，否则的话下次还是会遍历到相同的点

		if cur == end {
			break
		}

		for _, dir := range dirs { // 遍历cur的上 左 下 右 坐标; 从上开始以此遍历
			next := cur.add(dir) //next新发现节点: 当前的节点加上方向
			//比如当前节点是{1 1}，它的上面坐标计算就是{1+(-1) 1+0}，得到它的上面坐标{0,1}

			val, ok := next.at(maze) //next是point类型struct; at是它的方法; 用途: 取出next point坐标的的值
			if !ok || val == 1 {     //val == 1代表是墙
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 { //val != 0代表已经走过了
				continue
			}

			if next == start { //回到原点了
				continue
			}

			//如果新发现的节点都满足条件后，需要做两件事情
			//1. 把下一步的steps填进去，比如当前节点是1， 那么下一步根据坐标+1，变成2
			//2. 将新发现的next坐标，写到队列里面
			curSteps, _ := cur.at(steps) //先获取当前步数
			steps[next.i][next.j] =      //在去加1
				curSteps + 1

			//将next新发现的坐标压入队列, 等待下一次遍历
			Q = append(Q, next)
		}
	}

	return steps
}

func main() {
	maze := readMaze("D:\\代码\\Google资深工程师深度讲解Go语言\\coding-180\\lang\\maze\\maze.in")

	/*
		打印生成的迷宫
		for _, row := range maze {
			for _, val := range row {
				fmt.Printf("%3d", val)  %3d就是按3位对齐,排版好看
			}
			fmt.Println() //换行
		}
	*/
	steps := walk(maze, point{0, 0},
		point{len(maze) - 1, len(maze[0]) - 1})

	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val) //%3d 按3位对齐
		}
		fmt.Println()
	}
	/*
		生成后的steps，从0开始走，走到13的路径
		    0  0  4  5  6
			1  2  3  0  7
			2  0  4  0  8
			0  0  0 10  9
			0  0 12 11  0
			0  0 13 12 13
	*/
	// TODO: construct path from steps
}
