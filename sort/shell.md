#希尔排序
基本思想：在要排序的一组数中，根据某一增量分为若干子序列，并对子序列分别进行插入排序。然后逐渐将增量减小,并重复上述过程。直至增量为1,此时数据序列基本有序,最后进行插入排序。具体如下图所示： 
<img src="https://github.com/linpeixun/arithmetic-go/master/sort/shell.jpg"/>