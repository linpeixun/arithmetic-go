#选择算法
**基本思想：**在长度为N的无序数组中，第一次遍历n-1个数，找到最小的数值与第一个元素交换，第二次遍历n-2个数，找到最小的数值与第二个元素交换。。。第n-1次遍历，找到最小的数值与第n-1个元素交换，排序完成。具体如下图所示：
<img src="https://github.com/linpeixun/arithmetic-go/blob/master/sort/select.jpg?raw=true">