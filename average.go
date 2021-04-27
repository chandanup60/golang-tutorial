package main

import "fmt"

func main() {
/* int a;int b;int c;int d;int e;int t;double m;
a=Integer.parseInt(argu[0]);
b=Integer.parseInt(argu[1]);
c=Integer.parseInt(argu[2]);
d=Integer.parseInt(argu[3]);
e=Integer.parseInt(argu[4]);
t=a+b+c+d+e;
m=t/5;
System.out.println("Total="+t);
System.out.println("Average="+m);
} */

var a,b,c,d,e int;
fmt.Scanf("%d%d%d%d%d", &a, &b, &c, &d, &e);
var t int=a+b+c+d+e;
/* var m double;
m=t/5; */
fmt.Println(t);
//fmt.Println(m);

}